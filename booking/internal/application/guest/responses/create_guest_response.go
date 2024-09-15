package responses

import (
	"github.com/loxt/hotel-booking/booking/internal/application/guest/dto"
	"github.com/loxt/hotel-booking/booking/internal/application/shared"
)

type CreateGuestResponse struct {
	shared.Response
	GuestDTO dto.GuestDTO
}
