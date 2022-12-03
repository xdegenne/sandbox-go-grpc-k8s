package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = createRootCmd()

func createRootCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:           "hello",
		Short:         "hello - bla bla",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	return cmd
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nWhoops. There was an error while executing your CLI :\n=> %s\n", err)
		os.Exit(1)
	}
}
