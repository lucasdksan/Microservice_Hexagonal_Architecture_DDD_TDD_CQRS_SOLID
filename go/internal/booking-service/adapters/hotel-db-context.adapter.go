package adapters

import (
	"book/internal/booking-service/core/domain/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type HotelDbContext struct {
	DB *gorm.DB
}

func NewHotelDbContext(dsn string) (*HotelDbContext, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("postgres opening error: %v", err)
		return nil, err
	}

	db.AutoMigrate(&entities.Guest{}, &entities.Booking{}, &entities.Room{})

	return &HotelDbContext{DB: db}, nil
}
