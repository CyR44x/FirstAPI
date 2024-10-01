package main

import (
	"encoding/json"
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
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message requestBody
	json.NewDecoder(r.Body).Decode(&message)
	messageMessage := Message{Text: message.Message}
	DB.Create(&messageMessage)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var id, updateMessage requestBody
	json.NewDecoder(r.Body).Decode(&updateMessage)
	var message Message
	DB.Take(&message, id)
	message.Text = updateMessage.Message
	DB.Save(&message)
	json.NewEncoder(w).Encode(&message)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	//var id requestBody
	//json.NewDecoder(r.Body).Decode(&id)
	DB.Delete(&message, "ID")
	json.NewEncoder(w).Encode(&message)

}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	router.HandleFunc("/api/messages", UpdateMessage).Methods("PATCH")
	router.HandleFunc("/api/messages", DeleteMessage).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
