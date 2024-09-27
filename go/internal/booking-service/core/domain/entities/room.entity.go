package entities

type Room struct {
	Id            int
	Name          string
	Level         int
	InMaintenance bool
}

func (r Room) isAvailable() bool {
	if !r.InMaintenance || r.hasGuest() {
		return false
	}
	return true
}

func (r Room) hasGuest() bool {
	return true
}
