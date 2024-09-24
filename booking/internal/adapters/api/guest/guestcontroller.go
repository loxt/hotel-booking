package guest

import (
	"encoding/json"
	"fmt"
	"github.com/loxt/hotel-booking/booking/internal/application/guest/ports"
	"github.com/loxt/hotel-booking/booking/internal/application/guest/requests"
	"github.com/loxt/hotel-booking/booking/internal/application/shared"
	"net/http"
)

type Controller struct {
	manager ports.GuestManager
}

func NewGuestController(manager ports.GuestManager) *Controller {
	return &Controller{manager: manager}
}

func (gc *Controller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body requests.CreateGuestRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message": "invalid request payload"}`))
		return
	}

	res, err := gc.manager.CreateGuest(body)

	if err != nil || !res.Success {
		if res.ErrorCode == shared.InvalidEmail || res.ErrorCode == shared.InvalidPersonID || res.ErrorCode == shared.MissingRequiredInfo {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, res.Message)))
		return
	}

	if res.Success {
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(res.GuestDTO)
		return
	}
}
