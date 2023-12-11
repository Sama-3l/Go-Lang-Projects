package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func start(w http.ResponseWriter, req *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", start)
}
