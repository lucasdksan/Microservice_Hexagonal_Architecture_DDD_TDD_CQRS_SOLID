package main

import (
	"book/internal/booking-service/adapters"
	"book/internal/shared/infrastructure/config"
	"fmt"
	"log"
)

func main() {
	config.LoadingEnv()
	db, err := adapters.NewHotelDbContext(config.Bank_connection_string)

	if err != nil {
		log.Fatal()
	}

	fmt.Print(db)
}
