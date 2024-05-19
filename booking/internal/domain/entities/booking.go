package entities

import (
	"github.com/loxt/hotel-booking/booking/internal/domain/enums"
	"time"
)

type Booking struct {
	Id       string
	PlacedAt time.Time
	Start    time.Time
	End      time.Time
	status   enums.Status
}

func NewBooking() Booking {
	b := Booking{
		status: enums.Created,
	}

	return b
}

func (b *Booking) CurrentStatus() enums.Status {
	return b.status
}

func (b *Booking) ChangeState(action enums.Action) {
	if action == enums.Pay && b.status == enums.Created {
		b.status = enums.Paid
		return
	}

	if action == enums.Cancel && b.status == enums.Created {
		b.status = enums.Cancelled
		return
	}

	if action == enums.Finish && b.status == enums.Paid {
		b.status = enums.Finished
		return
	}

	if action == enums.Refund && b.status == enums.Paid {
		b.status = enums.Refunded
		return
	}

	if action == enums.Reopen && b.status == enums.Cancelled {
		b.status = enums.Created
		return
	}
}
