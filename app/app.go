package app

import (
	"fmt"
	"log"
	"net"

	proto "github.com/lutfipaper/module-proto/go/services/product"

	"google.golang.org/grpc"
)

var ServiceName = "nano-services"

func New() {

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	server := grpc.NewServer()
	GetConfig()

	m := NewModels()
	c := NewControllers(m)

	proto.RegisterServicesServer(server, c)

	//Run server
	fmt.Println("Run localhost:5001")
	if err := server.Serve(lis); err != nil {
		fmt.Println("err", err.Error())
	}
}
