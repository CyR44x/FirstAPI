package handlers

import (
	"RestApi/internal/userService"
	"RestApi/internal/web/users"
	"context"
	"fmt"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserService *userService.UserService
}

func NewUserHandler(userService *userService.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	userList, err := h.UserService.GetUsers()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := users.GetUsers200JSONResponse{}

	// Заполняем слайс response всеми юзерами из БД
	for _, usr := range userList {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	// Обращаемся к сервису и создаем юзера
	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.UserService.PostUser(userToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	// Получаем новое содержимое юзера из тела запроса
	newUser := request.Body
	if newUser == nil {
		// Возвращаем ошибку если новое поле юзера не указано
		return nil, fmt.Errorf("user content is required")
	}

	// Пытаемся обновить юзера по ID
	updatedUser, err := h.UserService.PatchUserByID(request.Id, userService.User{
		Email:    *newUser.Email,
		Password: *newUser.Password,
	})
	if err != nil {
		// Проверяем, если юзер не найден
		if err == gorm.ErrRecordNotFound {
			return users.PatchUsersId404Response{}, nil
		}
		// Возвращаем любую другую ошибку
		return nil, err
	}

	// Создаем ответ с обновленным сообщением
	response := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	// Возвращаем ответ
	return response, nil
}

func (h *UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.UserService.DeleteUserByID(request.Id)
	if err != nil {
		// Проверяем, если юзер не найден
		if err == gorm.ErrRecordNotFound {
			return users.DeleteUsersId404Response{}, nil
		}
		// Возвращаем любую другую ошибку
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}
