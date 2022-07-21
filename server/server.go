package main

import (
	"fmt"
	"io"
	"net/http"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// write `Hello` using `io.WriteString` function
	io.WriteString(res, "Hello")
	// write `World` using `fmt.Fprint` function
	fmt.Fprint(res, " World! ")
	// write `❤️` using simple `Write` call
	res.Write([]byte("❤️"))
}

func main() {

	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World! ❤️")
	})

	// handle `/hello/golang` route to `http.DefaultServeMux`
	http.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang! ❤️")
	})

	// listen and serve using `http.DefaultServeMux`
	http.ListenAndServe(":9000", nil)
}
