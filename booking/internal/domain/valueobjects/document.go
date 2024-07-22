package valueobjects

import "github.com/loxt/hotel-booking/booking/internal/domain/enums"

type Document struct {
	Value        string
	DocumentType enums.DocumentType
}
