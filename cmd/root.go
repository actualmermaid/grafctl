package cmd

import (
	"os"

	"github.com/dimitrovvlado/grafctl/grafana"

	"github.com/spf13/cobra"
)

type rootCmd struct {
	Verbose bool
	Output  string
}

//NewRootCmd creates the root command
func NewRootCmd(client *grafana.Client) *cobra.Command {
	i := &rootCmd{}

	rootCmd := &cobra.Command{
		Use:   "grafctl",
		Short: "Grafctl is command line tool for managing Grafana",
		Long:  `TODO`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&i.Verbose, "verbose", "v", false, "Verbose output")
	out := rootCmd.OutOrStdout()

	rootCmd.AddCommand(
		newVersionCmd(),
		newConfigCommand(out),
		newGetCmd(client, out),
		newCreateCmd(client, out),
		newDeletCmd(client, out))

	return rootCmd
}

// func newRootCmd(args []string) *cobra.Command {

// 	cmd := &cobra.Command{
// 		Use:          "grafctl",
// 		Short:        "The Grafana management tool.",
// 		Long:         globalUsage,
// 		SilenceUsage: true,
// 	}
// 	flags := cmd.PersistentFlags()
// 	flags.Parse(args)

// 	out := cmd.OutOrStdout()

// 	cmd.AddCommand(
// 		newGetCommand(out),
// 	)

// 	settings.Init()

// 	return cmd
// }
