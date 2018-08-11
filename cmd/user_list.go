package cmd

import (
	"fmt"

	"github.com/dimitrovvlado/grafctl/grafana"
	"github.com/gosuri/uitable"

	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

type usersCmd struct {
	client *grafana.Client
	output string
}

func newUsersCommand(client *grafana.Client) *cobra.Command {
	i := &usersCmd{
		client: client,
	}
	getUsersCmd := &cobra.Command{
		Use:     "users",
		Aliases: []string{"user"},
		Short:   "Display one or many users",
		Long:    `TODO`,
		RunE: func(cmd *cobra.Command, args []string) error {
			i.output = cmd.Flag("output").Value.String()
			return i.run()
		},
	}
	return getUsersCmd
}

// run creates a merge request
func (i *usersCmd) run() error {
	users, err := i.client.ListUsers()
	if err != nil {
		logrus.Fatal(err)
	}

	//TODO extract as flag
	var colWidth uint = 60
	formatter := func() string {
		table := uitable.New()
		table.MaxColWidth = colWidth
		table.AddRow("ID", "NAME", "LOGIN", "EMAIL")
		for _, lr := range users {
			table.AddRow(lr.ID, lr.Name, lr.Login, lr.Email)
		}
		return fmt.Sprintf("%s%s", table.String(), "\n")
	}

	result, err := formatResult(i.output, users, formatter)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf(result)

	return nil
}