package main

import (
	"RestApi/internal/database"
	"RestApi/internal/handlers"
	"RestApi/internal/messagesService"
	"RestApi/internal/userService"
	"RestApi/internal/web/messages"
	"RestApi/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&messagesService.Message{}); err != nil {
		log.Fatalf("Не удалось выполнить авто-миграцию: %v", err)
	}

	messagesRepo := messagesService.NewMessageRepository(database.DB)
	MessagesService := messagesService.NewService(messagesRepo)
	messagesHandler := handlers.NewHandler(MessagesService)

	userRepo := userService.NewUserRepository(database.DB)
	UserService := userService.NewService(userRepo)
	userHandlers := handlers.NewUserHandler(UserService)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictMessageHandler := messages.NewStrictHandler(messagesHandler, nil) // тут будет ошибка
	messages.RegisterHandlers(e, strictMessageHandler)

	// Регистрируем хендлеры пользователей
	strictUserHandler := users.NewStrictHandler(userHandlers, nil)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
