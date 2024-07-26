package responses

import (
	"github.com/loxt/hotel-booking/booking/internal/application"
	"github.com/loxt/hotel-booking/booking/internal/application/guest/dto"
)

type CreateGuestResponse struct {
	application.Response
	GuestDTO dto.GuestDTO
}
