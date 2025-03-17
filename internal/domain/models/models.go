package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UID          uuid.UUID `json:"uuid"`
	Username     string    `json:"username,omitempty" validate:"required"`
	Password     string    `json:"pass" validate:"required,min=8"`
	Email        string    `json:"email" validate:"required,email"`
	RegisterDate time.Time `json:"registerDate,omitempty"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"pass" validate:"required,min=8"`
}

type Product struct {
	ProductID   uuid.UUID `json:"uuid"`
	Name        string    `json:"name,omitempty" validate:"required"`
	Description string    `json:"description,omitempty" validate:"required"`
	Price       float64   `json:"price,omitempty" validate:"required"`
	Quantity    int       `json:"quantity,omitempty" validate:"required"`
}

type Purchase struct {
	PurchaseID uuid.UUID `json:"uuid"`
	UID        uuid.UUID `json:"uuid"`
	ProductID  uuid.UUID `json:"uuid"`
	Quantity   int       `json:"quantity,omitempty" validate:"required"`
	Timestamp  string    `json:"timestamp,omitempty" validate:"required"`
}
