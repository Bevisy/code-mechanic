package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HS.Hello", "world", &reply)
	if err != nil {
		log.Fatal("call failed:", err)
	}

	fmt.Println(reply)
}
