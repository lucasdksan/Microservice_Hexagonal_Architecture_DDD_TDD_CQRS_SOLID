package ports

import "book/internal/booking-service/core/domain/entities"

type IGuestRepository interface {
	Get(id uint) entities.Guest
	Create(guest entities.Guest) error
}
