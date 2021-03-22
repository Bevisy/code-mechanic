package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (p *HelloService) Hello(req *String, reply *String) error {
	reply.Value = "hello " + req.GetValue()
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		rpc.ServeConn(conn)
	}
}
