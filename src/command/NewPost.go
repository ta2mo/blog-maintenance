package command

import (
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

func NewPost(context *cli.Context) error {
	now := time.Now()
	title := context.Args().Get(1)
	if len(title) == 0 {
		return cli.NewExitError("Please input title.", 1)
	}
	categories := context.Args().Get(2)

	m := map[string]string{
		"title":      title,
		"categories": categories,
		"date":       now.Format("2006-01-02 03:04:05"),
	}
	renderNewPost(m, "post/"+NewPostTemplate, PostDir+now.Format("2006-01-02-")+"new"+MarkdownExt)
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
