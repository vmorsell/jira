package main

import (
	"fmt"

	"github.com/vmorsell/jira-cli/jira"

	cli "github.com/urfave/cli/v2"
)

var getIssue = &cli.Command{
	Name:    "issue",
	Aliases: []string{"i"},
	Action: func(c *cli.Context) error {
		if c.Args().Len() < 1 {
			return fmt.Errorf("missing issue ID or key")
		}

		idOrKey := c.Args().Get(0)

		client := jira.New().WithCredsFromFile()
		res, err := client.Issue(idOrKey)
		if err != nil {
			return err
		}
		fmt.Println(res)

		return nil
	},
}
