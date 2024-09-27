package ports

import "book/internal/booking-service/core/domain/entities"

type IGuestRepository interface {
	Get(id int) entities.Guest
	Save(guest entities.Guest) int
}
