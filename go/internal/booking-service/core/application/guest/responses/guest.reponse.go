package responses

import (
	"book/internal/booking-service/core/application/guest/dtos"
	"book/internal/shared/application/responses"
)

type GuestResponse struct {
	responses.Response
	Data dtos.GuestDTO
}
