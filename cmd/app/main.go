package main

import (
	"RestApi/internal/database"
	"RestApi/internal/handlers"
	"RestApi/internal/messagesService"
	"RestApi/internal/web/messages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := messages.NewStrictHandler(handler, nil) // тут будет ошибка
	messages.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}

//router := mux.NewRouter()
//router.HandleFunc("/api/post", handler.CreateMessageHandler).Methods("POST")
//router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
//router.HandleFunc("/api/patch/{id}", handler.UpdateMessageHandler).Methods("PATCH")
//router.HandleFunc("/api/delete/{id}", handler.DeleteMessageHandler).Methods("DELETE")

//http.ListenAndServe(":8080", router)
//}
