package repositories

import (
	"book/internal/booking-service/core/domain/entities"

	"gorm.io/gorm"
)

type GuestRepository struct {
	db *gorm.DB
}

func NewGuestRepository(db *gorm.DB) *GuestRepository {
	return &GuestRepository{
		db: db,
	}
}

func (gr *GuestRepository) Create(guest entities.Guest) error {
	if err := gr.db.Create(guest).Error; err != nil {
		return err
	}

	return nil

}

func (gr *GuestRepository) Get(id uint) entities.Guest {
	return entities.Guest{}
}
