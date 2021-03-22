package main

import (
	"log"
	"net"
	"net/rpc"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(req string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HellService struct{}

func (p *HellService) Hello(req string, reply *string) error {
	*reply = "hello " + req
	return nil
}

func main() {
	RegisterHelloService(new(HellService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpc.ServeConn(conn)
	}
}
