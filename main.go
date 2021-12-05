package main

import (
	"fmt"
	"log"
	"net"

	"github.com/imhshekhar47/go-grpc-pi/core"

	"google.golang.org/grpc"
)

var (
	appConfig = core.GetAppConfig()
)

func main() {
	log.Println("entry: main")
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", appConfig.Server.Port))
	if err != nil {
		log.Fatalf("Failed to launch server: %v", err)
	}

	server := grpc.NewServer()

	log.Printf("Launching server on 0.0.0.0:%v\n", appConfig.Server.Port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Filed to launch server: %v", err)
	}

	log.Println("exit: main")
}
