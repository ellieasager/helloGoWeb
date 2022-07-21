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

	// create a new `ServeMux`
	mux := http.NewServeMux()

	// handle `/` route
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World! ❤️")
	})

	// handle `/hello/golang` route
	mux.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang! ❤️")
	})

	// listen and serve using `ServeMux`
	http.ListenAndServe(":9000", mux)
}
