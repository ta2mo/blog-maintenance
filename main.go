package main

import (
	"github.com/ta2mo/blog-maintenance/internal/option"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "blog generator"
	app.Usage = "generate blog post from markdown file."
	app.Version = "1.0.0"

	app.Action = func(context *cli.Context) error {
		switch context.Args().Get(0) {
		case "convert", "c":
			if err := option.Convert(context); err != nil {
				return err
			}
		case "new", "n":
			if err := option.New(context); err != nil {
				return err
			}
		default:
			cli.ShowAppHelp(context)
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose, vvv",
			Usage: "print detail log",
		},
	}

	app.Run(os.Args) // 忘れない！（よくやらかす）
}
