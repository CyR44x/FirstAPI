package main

import (
	"RestApi/internal/database"
	"RestApi/internal/handlers"
	"RestApi/internal/messagesService"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/post", handler.CreateMessageHandler).Methods("POST")
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/patch/{id}", handler.UpdateMessageHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteMessageHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
