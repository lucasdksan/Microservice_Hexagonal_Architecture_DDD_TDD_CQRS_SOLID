package ports

import (
	"book/internal/booking-service/core/application/guest/requests"
	"book/internal/booking-service/core/application/guest/responses"
)

type IGuestManager interface {
	CreateGuest(request requests.GuestRequest) responses.GuestResponse
}
