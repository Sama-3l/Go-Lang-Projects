package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func createPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := ":" + port

	return addr
}

func hello(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(writer, "404", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(writer, "Wrong method", http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, "<h1>Hello, World!</h1>")
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)
	port := createPort()

	fmt.Printf("Server at port %v", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello World")
}
