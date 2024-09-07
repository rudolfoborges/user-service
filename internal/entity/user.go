package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Password  string
	AvatarUrl *string `db:"avatar_url"`
	Active    bool
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUser(name, email, password string) *User {
	id, _ := uuid.NewV7()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  string(hashedPassword),
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (e *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password))
}
