package application

import (
	"github.com/loxt/hotel-booking/booking/internal/application/guest/requests"
	"github.com/loxt/hotel-booking/booking/internal/application/guest/responses"
	"github.com/loxt/hotel-booking/booking/internal/domain/ports"
)

type GuestManager struct {
	repository ports.GuestRepository
}

func NewGuestManager(guestManager ports.GuestRepository) *GuestManager {
	return &GuestManager{
		repository: guestManager,
	}
}

func (gm *GuestManager) CreateGuest(request requests.CreateGuestRequest) (responses.CreateGuestResponse, error) {
	guest := request.ToEntity()
	err := gm.repository.Create(guest)

	if err != nil {
		return responses.CreateGuestResponse{
			Response: Response{
				Success:   false,
				Message:   err.Error(),
				ErrorCode: COULD_NOT_STORE_DATA,
			},
			GuestDTO: request.GuestDTO,
		}, err
	}

	return responses.CreateGuestResponse{
		Response: Response{
			Success: true,
		},
		GuestDTO: request.GuestDTO,
	}, nil
}
