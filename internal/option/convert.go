package option

import (
	"bufio"
	"github.com/microcosm-cc/bluemonday"
	_ "github.com/russross/blackfriday"
	"github.com/shurcooL/github_flavored_markdown"
	"github.com/ta2mo/blog-maintenance/internal/model"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"
	"unsafe"
)

const (
	vueExt           = ".vue"
	postDir          = "./post"
	templateDir      = "./internal/template/"
	nuxtPageDir      = "./nuxt-template/pages/post/"
	nuxtComponentDir = "./nuxt-template/components/"

	separatorLine = "^-*-$"

	SidebarDir    = "sidebar/"
	RecentPostVue = "RecentPost.vue"
	PostListVue   = "PostList.vue"
	CategoryVue   = "Category.vue"

	otherCategory = "その他"
)

var reg = regexp.MustCompile(separatorLine)

// Convert convert from markdown files to vue page conponents.
func Convert(context *cli.Context) error {
	posts, err := MappingFromFiles(postDir)
	if err != nil {
		return nil
	}

	sort.Slice(posts, func(i int, j int) bool {
		return posts[i].Header.Date.After(posts[j].Header.Date)
	})

	postTemplate := template.Must(template.ParseFiles(templateDir + "post.vue"))
	for i, post := range posts {
		f, err := os.Create(nuxtPageDir + post.FileName + vueExt)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		var olderPost *model.Post
		if len(posts) > i+1 {
			olderPost = &posts[i+1]
		}

		var newerPost *model.Post
		if i > 0 {
			newerPost = &posts[i-1]
		}

		params := map[string]interface{}{
			"post":      post,
			"newerPost": newerPost,
			"olderPost": olderPost,
		}

		postTemplate.Execute(f, params)
	}

	renderComponent(posts[:5], PostListVue, nuxtComponentDir+PostListVue)
	renderSidebar(posts[:5], RecentPostVue)
	renderCategory(posts, CategoryVue)

	return nil
}

// MappingFromFiles mapping post file to object.
func MappingFromFiles(targetDir string) ([]model.Post, error) {
	fileNames, err := ls(targetDir)
	if err != nil {
		return nil, err
	}

	var files []os.File
	var errors []error
	for _, name := range fileNames {
		func() {
			file, err := os.Open(targetDir + "/" + name)

			if err != nil {
				errors = append(errors, err)
			} else {
				files = append(files, *file)
			}
		}()
	}

	var posts []model.Post
	for _, file := range files {
		func() {
			post := parseFile(&file)
			posts = append(posts, *post)
		}()
		defer file.Close()
	}
	return posts, nil
}

func renderCategory(posts []model.Post, templateFileName string) {
	template := template.Must(template.ParseFiles(templateDir + templateFileName))
	f, err := os.Create(nuxtComponentDir + SidebarDir + templateFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := map[string][]model.Post{}
	for _, post := range posts {
		for _, category := range post.Header.Categories {
			if category == "" {
				m[otherCategory] = append(m[category], post)
			} else {
				m[category] = append(m[category], post)
			}
		}
	}

	template.Execute(f, m)
}

func renderSidebar(posts []model.Post, templateFileName string) {
	renderComponent(posts, templateFileName, nuxtComponentDir+SidebarDir+templateFileName)
}

func renderComponent(posts []model.Post, templateFileName string, outputFilePath string) {
	template := template.Must(template.ParseFiles(templateDir + templateFileName))
	f, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	template.Execute(f, posts)
}

func ls(dirName string) ([]string, error) {
	var fileNames []string
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileNames = append(fileNames, file.Name())
	}

	return fileNames, nil
}

func parseFile(file *os.File) *model.Post {
	post := model.Post{}
	s := bufio.NewScanner(file)

	rep := regexp.MustCompile(`.markdown$|.md$`)
	e := filepath.Base(rep.ReplaceAllString(file.Name(), ""))
	post.FileName = e
	separatorCount := 0
	for s.Scan() {
		line := s.Text()
		switch {
		case reg.MatchString(line):
			separatorCount++
		case separatorCount < 2:
			parseHeader(&post.Header, line)
		case 2 <= separatorCount:
			post.Content = post.Content + line + "\r\n"
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	if len(post.Content) > 0 {
		output := github_flavored_markdown.Markdown([]byte(post.Content))
		post.Content = *(*string)(unsafe.Pointer(&output))
		p := bluemonday.StrictPolicy().Sanitize(post.Content)
		post.RowContent = p
	}

	return &post
}

func parseHeader(header *model.PostHeader, line string) {
	l := strings.SplitN(line, ":", 2)
	switch {
	case l[0] == "layout":
		header.Layout = strings.TrimLeft(l[1], " ")
	case l[0] == "title":
		header.Title = strings.Trim(strings.TrimSpace(l[1]), "\"")
	case l[0] == "date":
		d, _ := time.Parse("2006-01-02 15:04:05 -0700", strings.TrimSpace(l[1]))
		header.Date = d
	case l[0] == "comments":
		header.Comments = strings.TrimSpace(l[1])
	case l[0] == "categories":
		header.Categories = parseListFormat(l[1])
	case l[0] == "tags":
		header.Tags = parseListFormat(l[1])
	}
}

func parseListFormat(str string) []string {
	var slice []string
	brackets := "(\\[|\\])"
	rep := regexp.MustCompile(brackets)
	s := rep.ReplaceAllString(str, " ")
	for _, str := range strings.Split(s, ",") {
		elm := strings.Trim(strings.TrimSpace(str), "\"")
		slice = append(slice, elm)
	}
	return slice
}
