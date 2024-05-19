package enums

type Status string

const (
	Created   Status = "created"
	Confirmed Status = "confirmed"
	Cancelled Status = "cancelled"
	Paid      Status = "paid"
	Finished  Status = "finished"
	Refunded  Status = "refunded"
)
