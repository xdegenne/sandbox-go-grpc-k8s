package cmd

import (
	server "example.com/hello/pkg/server"
	"github.com/spf13/cobra"
)

var serverCmd = createServerCmd()

func createServerCmd() *cobra.Command {

	config := server.ServerConfig{}

	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Launch the server",
		Run: func(cmd *cobra.Command, args []string) {
			server := server.Server{Config: config}
			server.Start()
		},
	}

	serverCmd.PersistentFlags().StringVar(&config.Address, "address", ":5555", "The address of the TCP server")

	return serverCmd
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
