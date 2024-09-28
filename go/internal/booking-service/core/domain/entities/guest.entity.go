package entities

import (
	valueobjects "book/internal/booking-service/core/domain/value-objects"
)

type Guest struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Surname  string
	Email    string
	Document valueobjects.Person `gorm:"embedded"`
}
