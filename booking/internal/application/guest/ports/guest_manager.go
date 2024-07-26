package ports

import (
	"github.com/loxt/hotel-booking/booking/internal/application/guest/requests"
	"github.com/loxt/hotel-booking/booking/internal/application/guest/responses"
)

type GuestManager interface {
	CreateGuest(request requests.CreateGuestRequest) (responses.CreateGuestResponse, error)
}
