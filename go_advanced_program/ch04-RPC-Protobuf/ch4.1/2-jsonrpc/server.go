package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloSVC interface {
	Hello(req string, reply *string) error
}

type HelloSVCServer struct {
}

func (p *HelloSVCServer) Hello(req string, reply *string) error {
	*reply = "hello: " + req
	return nil
}

func main() {
	rpc.RegisterName("HelloSVC", new(HelloSVCServer))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
