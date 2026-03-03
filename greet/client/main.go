package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/Omar-Sa6ry/grpc-go/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	var opts []grpc.DialOption
	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Failed to load TLS credentials: %v", err)
		}
		
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	doGreet(client)
	// doGreetManyTimes(client)
	// doLongGreet(client)
	// doGreetEveryone(client)
	// doGreetWithDeadline(client, 5*time.Second)
	// doGreetWithDeadline(client, 1*time.Second)
}