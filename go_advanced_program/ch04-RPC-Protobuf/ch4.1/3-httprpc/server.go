package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloSVC struct {
}

func (p *HelloSVC) Hello(req string, reply *string) error {
	*reply = "hello " + req
	return nil
}

func main() {
	rpc.RegisterName("HelloSVC", new(HelloSVC))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}

// 客户端请求： curl localhost:1234/jsonrpc -X POST --data '{"method":"HelloSVC.Hello","params":["world"],"id":0}'
