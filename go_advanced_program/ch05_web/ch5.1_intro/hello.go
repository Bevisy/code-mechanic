package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func echo(wr http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte("there is a err."))
		return
	}

	msg = append(msg, []byte("hello")...)

	c, err := wr.Write(msg)
	if err != nil || c != len(msg) {
		log.Println(err, "write len:", c)
	}
}

func main() {
	http.HandleFunc("/", echo)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
