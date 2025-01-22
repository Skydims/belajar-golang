package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/dimas", handler.DimasHandler)

	log.Println("mulai di web 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
