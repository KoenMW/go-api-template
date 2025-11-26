package main

import (
	helloworld "go-api/adaptors/hello_world"
	"go-api/adaptors/rest"
	"net/http"
)

func AddEndpoint(endpoint string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(endpoint, handler)
}

func main() {

	producer := &helloworld.HelloWorldProducer{}
	handler := &rest.Handler{Producer: producer}
	AddEndpoint("/hello", handler.Messages)

	http.ListenAndServe(":8080", nil)
}
