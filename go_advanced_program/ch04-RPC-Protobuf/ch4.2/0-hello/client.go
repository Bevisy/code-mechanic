package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	var reply String
	req := String{
		Value: "world",
	}

	// 方法参数reply 传递指针报错： reading body gob: DecodeValue of unassignable value
	err = client.Call("HelloService.Hello", &req, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

}
