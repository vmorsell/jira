package main

import (
	"fmt"

	"github.com/vmorsell/jira-cli/authstore"

	cli "github.com/urfave/cli/v2"
)

var auth = &cli.Command{
	Name:    "auth",
	Aliases: []string{"a"},
	Flags: []cli.Flag{
		&cli.StringFlag{Name: "tenant"},
		&cli.StringFlag{Name: "email"},
		&cli.StringFlag{Name: "token"},
	},
	Action: func(c *cli.Context) error {
		s := authstore.New()

		var creds *authstore.Credentials
		creds, err := s.Read()
		if err != nil {
			creds = &authstore.Credentials{}
		}

		if v := c.String("tenant"); v != "" {
			creds.Tenant = v
		}
		if v := c.String("email"); v != "" {
			creds.Email = v
		}
		if v := c.String("token"); v != "" {
			creds.Token = v
		}

		if err := s.Write(creds); err != nil {
			return fmt.Errorf("write: %w", err)
		}
		return nil
	},
}
