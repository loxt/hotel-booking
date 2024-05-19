package entities

type Room struct {
	Id            string
	Name          string
	Level         int
	InMaintenance bool
}

func (r *Room) IsAvailable() bool {
	return !r.InMaintenance || !r.HasGuests()
}

func (r *Room) HasGuests() bool {
	return true
}
