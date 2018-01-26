package main

import (
	"fmt"
	"net/http"
)

//type handler func(http.ResponseWriter, *http.Request) func(w http.ResponseWriter, r *http.Request)

func main() {
	http.HandleFunc("/test", httpHandler("GET", helloHandler))
	http.ListenAndServe(":8080", nil)
}

func httpHandler (method string, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func (res http.ResponseWriter, req *http.Request) {
		if req.Method == method {
			handler(res, req)
		}
	}
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "ovo je neki odgovor")
}
