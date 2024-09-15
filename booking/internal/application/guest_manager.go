package application

import (
	"github.com/loxt/hotel-booking/booking/internal/application/guest/requests"
	"github.com/loxt/hotel-booking/booking/internal/application/guest/responses"
	"github.com/loxt/hotel-booking/booking/internal/application/shared"
	"github.com/loxt/hotel-booking/booking/internal/domain/ports"
)

type GuestManager struct {
	service ports.GuestService
}

func NewGuestManager(service ports.GuestService) *GuestManager {
	return &GuestManager{service}
}

func (gm *GuestManager) CreateGuest(request requests.CreateGuestRequest) (responses.CreateGuestResponse, error) {
	guest := request.ToEntity()
	err := gm.service.Save(guest)

	if err != nil {
		return responses.CreateGuestResponse{
			Response: shared.Response{
				Success:   false,
				Message:   err.Error(),
				ErrorCode: shared.CouldNotStoreData,
			},
			GuestDTO: request.GuestDTO,
		}, err
	}

	return responses.CreateGuestResponse{
		Response: shared.Response{
			Success: true,
		},
		GuestDTO: request.GuestDTO.FromEntity(guest),
	}, nil
}
