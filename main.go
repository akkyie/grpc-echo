package main

import (
	"fmt"
	"log"
	"net"

	"github.com/akkyie/grpc-echo/echo"
	"github.com/akkyie/grpc-echo/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	server, err := server.NewServer()
	if err != nil {
		log.Fatalln(err)
	}

	echo.RegisterEchoServer(grpcServer, server)
	reflection.Register(grpcServer)

	fmt.Printf("listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
