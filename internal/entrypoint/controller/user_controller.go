package controller

import (
	"encoding/json"
	"net/http"

	"github.com/amandakeren/user-service/internal/helper"
	"github.com/amandakeren/user-service/internal/service"
	"github.com/go-chi/chi/v5"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Routes(mux *chi.Mux) {
	mux.Handle("POST /users", helper.HandleFunc(c.create))
}

func (c *UserController) create(w http.ResponseWriter, r *http.Request) error {
	var requestBody CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return err
	}

	err = c.userService.Create(r.Context(), service.CreateUserInput{
		Name:     requestBody.Name,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)

	return nil
}
