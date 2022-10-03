package main

import (
	"fmt"
	"io"
	"net/http"
)




func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	if method := r.Method; method=="POST"{
		io.WriteString(w, "<h1>Unauthorized method</h1>")
		return
	}
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
