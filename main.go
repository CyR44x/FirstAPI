package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type requestBody struct {
	ID      uint   `json:"id"`
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
	DB.Create(&Message{Text: message.Message})
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	var message requestBody
	json.NewDecoder(r.Body).Decode(&message)
	DB.Model(&Message{}).Where("id = ?", &message.ID).Update("text", message.Message)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var message requestBody
	json.NewDecoder(r.Body).Decode(&message)
	// Удаляем запись по id
	DB.Model(&Message{}).Where("id = ?", message.ID).Delete(&Message{})
	//DB.Model(&Message{}).Where("id = ?", message.ID).Update("deleted_at", time.Now())
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
