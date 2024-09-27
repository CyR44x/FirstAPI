package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type requestBody struct {
	Message string `json:"message"`
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	var message []Message
	DB.Find(&message)
	json.NewEncoder(w).Encode(&message)
	fmt.Fprint(w, message)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message requestBody
	json.NewDecoder(r.Body).Decode(&message)
	messageMessage := Message{Text: message.Message}
	DB.Create(&messageMessage)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)
}
