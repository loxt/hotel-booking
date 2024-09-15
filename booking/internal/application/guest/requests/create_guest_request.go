package requests

import "github.com/loxt/hotel-booking/booking/internal/application/guest/dto"

type CreateGuestRequest struct {
	dto.GuestDTO
	Id string `json:"-"`
}
