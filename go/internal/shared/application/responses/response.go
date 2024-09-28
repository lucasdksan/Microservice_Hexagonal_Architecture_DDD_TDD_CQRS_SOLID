package responses

type ErrorCodes int

const (
	NOT_FOUND           ErrorCodes = 1
	COULDNOT_STORE_DATA ErrorCodes = 2
)

type Response struct {
	Success   bool
	Message   string
	ErrorCode ErrorCodes
}
