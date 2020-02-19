package option

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"text/template"
	"time"
)

const (
	MarkdownExt     = ".md"
	NewPostTemplate = "new.md"
	PostDir         = "post/"
)

func New(_ *cli.Context) error {
	now := time.Now()
	var title string

	fmt.Println("Please input post information.")
	fmt.Print("post title       : ")
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		title = sc.Text()
		break
	}

	fmt.Print("post categories  : ")
	var categories string
	for sc.Scan() {
		categories = sc.Text()
		break
	}

	m := map[string]string{
		"title":      title,
		"categories": categories,
		"date":       now.Format("2006-01-02 15:04:05"),
	}
	fileName := PostDir + now.Format("2006-01-02-") + "new" + MarkdownExt
	renderNewPost(m, "post/"+NewPostTemplate, fileName)
	fmt.Println("generete success : ./" + fileName)
	return nil
}

func renderNewPost(params map[string]string, templateFileName string, outputFilePath string) {
	template := template.Must(template.ParseFiles(templateDir + templateFileName))
	f, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	template.Execute(f, params)
}
