package tlsclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"example.com/hello/pkg/helloworld"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

type Client struct {
	Address    string
	client     helloworld.GreeterClient
	connection *grpc.ClientConn
}

func (c *Client) Connect() error {
	fmt.Println("Connecting")

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(tlsCredentials),
		//grpc.WithUnaryInterceptor(interceptor.Unary()),
		//grpc.WithStreamInterceptor(interceptor.Stream()),
	}

	// establish connection
	conn, err := grpc.Dial(c.Address, opts...)
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

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func (c *Client) SayHello(name string) error {
	in := &helloworld.HelloRequest{Name: name}
	reply, err := c.client.SayHello(context.Background(), in)
	if err != nil {
		fmt.Printf("Error: %v", err)

		return err
	}
	fmt.Printf("Response: ", reply.Message)
	return nil
}
