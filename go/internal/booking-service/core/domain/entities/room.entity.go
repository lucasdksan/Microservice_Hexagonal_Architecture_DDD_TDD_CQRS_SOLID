package entities

import valueobjects "book/internal/booking-service/core/domain/value-objects"

type Room struct {
	Id            uint `gorm:"primaryKey"`
	Name          string
	Level         int
	InMaintenance bool
	Price         valueobjects.Price `gorm:"embedded"`
}

func (r Room) IsAvailable() bool {
	if !r.InMaintenance || r.HasGuest() {
		return false
	}
	return true
}

func (r Room) HasGuest() bool {
	return true
}
