package main

import "example.com/hello/pkg/cmd"

func main() {
	//s := server.Server{}
	//s.Start()

	//c := client.Client{Address: "hello.example.com:443"}
	////c := client.Client{Address: "localhost:5555"}
	//c.Connect()
	//c.SayHello("tonton")
	//c.Disconnect()

	cmd.Execute()
}
