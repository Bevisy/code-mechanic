package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

// 必须是公开的方法
// 方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

func main() {
	rpc.RegisterName("HS", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
