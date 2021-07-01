package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "Jira CLI",
		Commands: []*cli.Command{
			{
				Name: "get",
				Subcommands: []*cli.Command{
					getIssue,
				},
			},
			auth,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
