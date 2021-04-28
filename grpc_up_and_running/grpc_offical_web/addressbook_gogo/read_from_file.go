package main

import (
	"fmt"
	"io/ioutil"
	"os"

	pb "github.com/Bevisy/code-mechanic/grpc_up_and_running/grpc_offical_web/addressbook/tutorialpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("HELP: ./read_from_file filename")
		os.Exit(1)
	}

	fname := os.Args[1]
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	book := &pb.AddressBook{}
	err = proto.Unmarshal(in, book)
	if err != nil {
		fmt.Println("proto Unmarshal failed")
	}

	fmt.Println(book.String()) // people:{name:"a" id:1 email:"1@qq.com" phones:{number:"123" type:HOME}} people:{name:"b" id:2 email:"2@qq.com" phones:{number:"123456"}}
}
