package enums

type Status int

const (
	Created Status = iota
	Paid
	Finished
	Canceled
	Refounded
)
