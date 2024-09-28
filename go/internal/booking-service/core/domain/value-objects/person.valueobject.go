package valueobjects

import "book/internal/booking-service/core/domain/enums"

type Person struct {
	IdNumber     string
	DocumentType enums.DocType
}
