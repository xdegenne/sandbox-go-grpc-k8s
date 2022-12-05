package cmd

import (
	"example.com/hello/pkg/client"
	"github.com/spf13/cobra"
)

var clientCmd = createClientCmd()

func createClientCmd() *cobra.Command {

	clientConfig := client.ClientConfig{}

	var clientCmd = &cobra.Command{
		Use:   "client",
		Short: "Launch the client",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			c := client.Client{Config: clientConfig}
			c.Connect()
			c.SayHello(args[0])
			c.Disconnect()

		},
	}

	clientCmd.PersistentFlags().StringVar(&clientConfig.Address, "address", "hello.example.com:443", "The address of the TCP server")
	clientCmd.PersistentFlags().StringVar(&clientConfig.CaCcert, "ca-cert", "", "The path of the CA certificate")
	clientCmd.PersistentFlags().BoolVar(&clientConfig.Tls, "tls", false, "Activate TLS")
	clientCmd.PersistentFlags().BoolVar(&clientConfig.Insecure, "insecure", false, "Trust all servers, even if self signed certs")
	return clientCmd
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
