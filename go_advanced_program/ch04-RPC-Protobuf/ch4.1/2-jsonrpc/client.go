package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net dial: ", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string

	err = client.Call("HelloSVC.Hello", "world", &reply)
	if err != nil {
		log.Fatal("call: ", err)
	}

	fmt.Println(reply)
}
