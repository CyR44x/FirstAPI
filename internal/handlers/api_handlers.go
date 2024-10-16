package handlers

import (
	"RestApi/internal/messagesService" // Импортируем наш сервис
	"RestApi/internal/web/messages"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type Handler struct {
	Service *messagesService.MessageService
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	// Получение всех сообщений из сервиса
	allMessages, err := h.Service.GetAllMessages()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := messages.GetMessages200JSONResponse{}

	// Заполняем слайс response всеми сообщениями из БД
	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	messageRequest := request.Body
	// Обращаемся к сервису и создаем сообщение
	messageToCreate := messagesService.Message{Text: *messageRequest.Message}
	createdMessage, err := h.Service.CreateMessage(messageToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *Handler) PatchMessagesId(_ context.Context, request messages.PatchMessagesIdRequestObject) (messages.PatchMessagesIdResponseObject, error) {
	// Получаем новое содержимое сообщения из тела запроса
	newMessage := request.Body.Message
	if newMessage == nil {
		// Возвращаем ошибку если новое сообщение не указано
		return nil, fmt.Errorf("message content is required")
	}

	// Пытаемся обновить сообщение по ID
	updatedMessage, err := h.Service.UpdateMessageByID(request.Id, messagesService.Message{Text: *newMessage})
	if err != nil {
		// Проверяем, если сообщение не найдено
		if err == gorm.ErrRecordNotFound {
			return messages.PatchMessagesId404Response{}, nil
		}
		// Возвращаем любую другую ошибку
		return nil, err
	}

	// Создаем ответ с обновленным сообщением
	response := messages.PatchMessagesId200JSONResponse{
		Id:      &updatedMessage.ID,
		Message: &updatedMessage.Text,
	}

	// Возвращаем ответ
	return response, nil
}

func (h *Handler) DeleteMessagesId(_ context.Context, request messages.DeleteMessagesIdRequestObject) (messages.DeleteMessagesIdResponseObject, error) {
	// Пытаемся удалить сообщение по ID
	err := h.Service.DeleteMessageByID(request.Id)
	if err != nil {
		// Проверяем, если сообщение не найдено
		if err == gorm.ErrRecordNotFound {
			return messages.DeleteMessagesId404Response{}, nil
		}
		// Возвращаем любую другую ошибку
		return nil, err
	}

	// Возвращаем 204 No Content в случае успеха
	return messages.DeleteMessagesId204Response{}, nil
}
