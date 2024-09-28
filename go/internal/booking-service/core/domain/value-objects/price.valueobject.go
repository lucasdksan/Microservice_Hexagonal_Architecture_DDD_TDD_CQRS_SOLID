package valueobjects

import "book/internal/booking-service/core/domain/enums"

type Price struct {
	Value    float32
	Currency enums.AcceptedCurrencies
}
