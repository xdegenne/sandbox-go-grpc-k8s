package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"example.com/hello/pkg/helloworld"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"log"
)

type Client struct {
	Config     ClientConfig
	client     helloworld.GreeterClient
	connection *grpc.ClientConn
}

type ClientConfig struct {
	Address  string
	Tls      bool
	CaCcert  string
	Insecure bool
}

func (c *Client) loadTLSCredentials() (credentials.TransportCredentials, error) {
	fmt.Println("Activating TLS")
	certPool := x509.NewCertPool()

	// Load certificate of the CA who signed server's certificate
	if c.Config.CaCcert != "" {
		fmt.Printf("Loading ca cart from %s", c.Config.CaCcert)
		pemServerCA, err := ioutil.ReadFile(c.Config.CaCcert)
		if err != nil {
			return nil, err
		}
		if !certPool.AppendCertsFromPEM(pemServerCA) {
			return nil, fmt.Errorf("Failed to add server CA's certificate")
		}
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}
	if c.Config.Insecure {
		fmt.Println("Skip verifying CA certs from server")
		config.InsecureSkipVerify = true
	}

	return credentials.NewTLS(config), nil
}

func (c *Client) Connect() error {
	fmt.Printf("Connecting to: %s\n", c.Config.Address)

	opts := []grpc.DialOption{}
	if c.Config.Tls {
		tlsCredentials, err := c.loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))
	} else {
		fmt.Println("No TLS activated")
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// establish connection
	conn, err := grpc.Dial(c.Config.Address, opts...)
	if err != nil {
		log.Fatalf("can not connect %v", err)
		return err
	}
	c.connection = conn
	// create client
	c.client = helloworld.NewGreeterClient(conn)
	return nil
}

func (c *Client) Disconnect() error {
	return c.connection.Close()
}

func (c *Client) SayHello(name string) error {
	in := &helloworld.HelloRequest{Name: name}
	reply, err := c.client.SayHello(context.Background(), in)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return err
	}
	fmt.Printf("Server response: %s", reply.Message)
	return nil
}
