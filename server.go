package proto

import (
	"google.golang.org/grpc"
	"net"
)

type ServerRpc struct {
	UnimplementedServerServer
}

func NewRpcServer() {
	grpcServer := grpc.NewServer()
	u := &ServerRpc{}
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	RegisterServerServer(grpcServer, u)
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
