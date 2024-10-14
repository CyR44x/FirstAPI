package handlers

import (
	"RestApi/internal/messagesService" // Импортируем наш сервис
	"RestApi/internal/web/messages"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *messagesService.MessageService
}

//func (h Handler) GetMessages(ctx context.Context, request messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
//TODO implement me
//panic("implement me")
//}

//func (h Handler) PostMessages(ctx context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
//TODO implement me
//panic("implement me")
//}

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

//func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
//messages, err := h.Service.GetAllMessages()
//if err != nil {
//http.Error(w, err.Error(), http.StatusInternalServerError)
//return
//}
//w.Header().Set("Content-Type", "application/json")
//json.NewEncoder(w).Encode(messages)
//}

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

//func (h *Handler) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
//var message messagesService.Message
//err := json.NewDecoder(r.Body).Decode(&message)
//if err != nil {
//http.Error(w, "Invalid input", http.StatusBadRequest)
//return
//}
//if message.Text == "" {
//http.Error(w, "Message text is required", http.StatusBadRequest)
//return
//}

//createdMessage, err := h.Service.CreateMessage(message)
//if err != nil {
//http.Error(w, err.Error(), http.StatusInternalServerError)
//return
//}
//log.Printf("Передаем в БД %v", createdMessage)

//w.Header().Set("Content-Type", "application/json")
//w.WriteHeader(http.StatusCreated)
//json.NewEncoder(w).Encode(createdMessage)
//}

func (h *Handler) UpdateMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var updatedMessageData messagesService.Message
	err = json.NewDecoder(r.Body).Decode(&updatedMessageData)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	updatedMessage, err := h.Service.UpdateMessageByID(id, updatedMessageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedMessage)
}

func (h *Handler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteMessageByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
