package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("entry: main")
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("Failed to launch server: %v", err)
	}

	server := grpc.NewServer()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Filed to launch server: %v", err)
	}

	log.Println("exit: main")
}
