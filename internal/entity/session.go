package entity

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Session struct {
	ID        int64
	UserID    uuid.UUID `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

func NewSession(userID uuid.UUID) *Session {
	return &Session{
		UserID:    userID,
		CreatedAt: time.Now(),
	}
}

func (e *Session) GenerateToken() (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = e.UserID
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	createdToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return createdToken, nil
}
