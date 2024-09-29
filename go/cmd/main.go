package main

import (
	"book/internal/booking-service/adapters"
	"book/internal/booking-service/adapters/repositories"
	"book/internal/booking-service/consumers/controllers"
	"book/internal/booking-service/core/application"
	"book/internal/shared/infrastructure/config"
	"book/internal/shared/infrastructure/router"
	"log"
)

func main() {
	if err := config.LoadingEnv(); err != nil {
		log.Fatal("fail in read file .env!")
	}

	db, err := adapters.NewHotelDbContext(config.Bank_connection_string)

	guestRepository := repositories.NewGuestRepository(db.DB)
	guestManager := application.NewGuestManager(guestRepository)
	guestController := controllers.NewGuestController(guestManager)

	if err != nil {
		log.Fatal()
	}

	router.Initialize(guestController)
}
