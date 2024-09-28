package ports

import "book/internal/booking-service/core/domain/entities"

type IGuestRepository interface {
	Get(id uint) entities.Guest
	Save(guest entities.Guest) uint
}
