package entities

import (
	valueobjects "book/internal/booking-service/core/domain/value-objects"
)

type Guest struct {
	Id       uint                `gorm:"primaryKey"`
	Name     string              `json:"name"`
	Surname  string              `json:"surname"`
	Email    string              `json:"email"`
	Document valueobjects.Person `gorm:"embedded"`
}
