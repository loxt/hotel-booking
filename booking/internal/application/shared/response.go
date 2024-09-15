package shared

type ErrorCodes string

const (
	NotFound          ErrorCodes = "NOT_FOUND"
	CouldNotStoreData ErrorCodes = "COULD_NOT_STORE_DATA"
)

type Response struct {
	Success   bool
	Message   string
	ErrorCode ErrorCodes
}
