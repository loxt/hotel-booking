package valueobjects

import "github.com/loxt/hotel-booking/booking/internal/domain/enums"

type Price struct {
	Value    float32
	Currency enums.AcceptedCurrencies
}
