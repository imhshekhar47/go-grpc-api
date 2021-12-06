package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"github.com/imhshekhar47/go-grpc-api/core"
	"github.com/imhshekhar47/go-grpc-api/pb/actuator"
	"google.golang.org/grpc"
)

var (
	appConfig          = core.GetAppConfig()
	actuatorServerImpl = actuator.NewServer(appConfig)
)

func main() {
	log.Println("entry: main")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", appConfig.Server.Port))
	if err != nil {
		log.Fatalf("Failed to launch server: %v", err)
	}

	grpcServer := grpc.NewServer()

	actuator.RegisterActuatorServiceServer(grpcServer, actuatorServerImpl)
	reflection.Register(grpcServer)

	log.Printf("Launching server on 0.0.0.0:%v\n", appConfig.Server.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Filed to launch server: %v\n", err)
	}

	log.Println("exit: main")
}
