package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unksato/mux-sample/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.IndexHandler)
	r.HandleFunc("/user/{username}", handler.UserHandler)
	http.ListenAndServe(":8080", r)
}
