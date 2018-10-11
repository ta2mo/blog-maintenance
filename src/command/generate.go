package command

import (
	"bufio"
	"github.com/microcosm-cc/bluemonday"
	_ "github.com/russross/blackfriday"
	"github.com/shurcooL/github_flavored_markdown"
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
	templateDir      = "./src/template/"
	nuxtPageDir      = "./nuxt-template/pages/post/"
	nuxtComponentDir = "./nuxt-template/components/"

	separatorLine = "^-*-$"
)

type Post struct {
	FileName string
	Header   PostHeader
	Content  string
}

type PostHeader struct {
	Layout     string
	Title      string
	Date       time.Time
	Comments   string
	Categories []string
	Tags       []string
}

var reg = regexp.MustCompile(separatorLine)

func Generate(context *cli.Context) error {
	fileNames, err := ls(postDir)
	if err != nil {
		return err
	}

	var files []os.File
	var errors []error
	for _, name := range fileNames {
		func() {
			file, err := os.Open(postDir + "/" + name)

			if err != nil {
				errors = append(errors, err)
			} else {
				files = append(files, *file)
			}
		}()
	}

	var posts []Post
	for _, file := range files {
		func() {
			post := parseFile(&file)
			posts = append(posts, *post)
		}()
		defer file.Close()
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

		var olderPost *Post
		if len(posts) > i+1 {
			olderPost = &posts[i+1]
		}

		var newerPost *Post
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

	postListTemplate := template.Must(template.ParseFiles(templateDir + "postList.vue"))
	f, err := os.Create(nuxtComponentDir + "postList.vue")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	postListTemplate.Execute(f, posts[:5])

	newPostsTemplate := template.Must(template.ParseFiles(templateDir + "NewPost.vue"))
	f, err = os.Create(nuxtComponentDir + "sidebar/NewPost.vue")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	newPostsTemplate.Execute(f, posts[:5])

	allPostsTemplate := template.Must(template.ParseFiles(templateDir + "AllPost.vue"))
	f, err = os.Create(nuxtComponentDir + "sidebar/AllPost.vue")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	allPostsTemplate.Execute(f, posts)

	return nil
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

func parseFile(file *os.File) *Post {
	post := Post{}
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
		output := bluemonday.UGCPolicy().SanitizeBytes(github_flavored_markdown.Markdown([]byte(post.Content)))
		post.Content = *(*string)(unsafe.Pointer(&output))
	}

	return &post
}

func parseHeader(header *PostHeader, line string) {
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
