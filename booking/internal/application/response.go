package application

type ErrorCodes string

const (
	NOT_FOUND            ErrorCodes = "NOT_FOUND"
	COULD_NOT_STORE_DATA ErrorCodes = "COULD_NOT_STORE_DATA"
)

type Response struct {
	Success   bool
	Message   string
	ErrorCode ErrorCodes
}
