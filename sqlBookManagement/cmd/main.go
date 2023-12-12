package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sqlBookManagement/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	port := createPort()
	fmt.Printf("%v", port)
	log.Fatal(http.ListenAndServe("localhost"+port, r))
}
