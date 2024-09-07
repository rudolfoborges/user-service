package controller

import (
	"encoding/json"
	"net/http"

	"github.com/amandakeren/user-service/internal/helper"
	"github.com/amandakeren/user-service/internal/service"
	"github.com/go-chi/chi/v5"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	SessionToken string `json:"session_token"`
}

type SessionController struct {
	sessionService *service.SessionService
}

func NewSessionController(sessionService *service.SessionService) *SessionController {
	return &SessionController{
		sessionService: sessionService,
	}
}

func (c *SessionController) Routes(mux *chi.Mux) {
	mux.Handle("POST /login", helper.HandleFunc(c.login))
}

func (c *SessionController) login(w http.ResponseWriter, r *http.Request) error {
	var requestBody LoginRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		return err
	}

	output, err := c.sessionService.Login(
		r.Context(), requestBody.Email, requestBody.Password,
	)

	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(LoginResponse{
		UserID:       output.UserID.String(),
		UserName:     output.UserName,
		SessionToken: output.SessionToken,
	})

	if err != nil {
		return err
	}

	return nil
}
