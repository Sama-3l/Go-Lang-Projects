package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"stocksAPI/routes"
)

func createPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := ":" + port

	return addr
}

func main() {
	r := routes.Router()
	port := createPort()
	fmt.Printf("Starting server at port%v\n", port)

	log.Fatal(http.ListenAndServe(port, r))
}
