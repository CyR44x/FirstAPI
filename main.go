package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var message requestBody

type requestBody struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %v", message)
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&message)
}

func main() {
	router := mux.NewRouter()
	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/hello", MessageHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
