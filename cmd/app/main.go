package main

import (
	"RestApi/internal/database"
	"RestApi/internal/handlers"
	"RestApi/internal/messagesService"
	"github.com/gorilla/mux"
	"net/http"
)

/*type requestBody struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	var message []Message
	DB.Find(&message)
	json.NewEncoder(w).Encode(&message)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message requestBody
	json.NewDecoder(r.Body).Decode(&message)
	DB.Create(&Message{Text: message.Message})
}

func UpdateMessageByID(w http.ResponseWriter, r *http.Request) {
	var message requestBody
	json.NewDecoder(r.Body).Decode(&message)
	DB.Model(&Message{}).Where("id = ?", &message.ID).Update("text", message.Message)
}

func DeleteMessageByID(w http.ResponseWriter, r *http.Request) {
	var message requestBody
	json.NewDecoder(r.Body).Decode(&message)
	DB.Model(&Message{}).Where("id = ?", message.ID).Delete(&Message{})

}*/

func main() {
	// Вызываем метод InitDB() из файла db.go
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	// Автоматическая миграция модели Message
	//DB.AutoMigrate(&Message{})
	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	//router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	//router.HandleFunc("/api/messages", GetAllMessages).Methods("GET")
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	//router.HandleFunc("/api/messages", UpdateMessageByID).Methods("PATCH")
	router.HandleFunc("/api/patch", handler.UpdateMessageHandler).Methods("PATCH")
	//router.HandleFunc("/api/messages", DeleteMessageByID).Methods("DELETE")
	router.HandleFunc("/api/delete", handler.DeleteMessageHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
