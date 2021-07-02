package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/vmorsell/jira/jira"

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
		if res.Fields == nil {
			fmt.Printf("%s (missing data)\n", res.Key)
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 10, 0, 3, ' ', 0)
		fmt.Fprintf(w, "%s\t%s\t(%s)\t[%s]\n", res.Key, res.Fields.Summary, res.Fields.Assignee.DisplayName, res.Fields.Status.Name)
		for _, s := range res.Fields.Subtasks {
			fmt.Fprintf(w, " - %s\t%s\t(%s)\t[%s]\n", s.Key, s.Fields.Summary, res.Fields.Assignee.DisplayName, res.Fields.Status.Name)
		}
		w.Flush()

		return nil
	},
}
