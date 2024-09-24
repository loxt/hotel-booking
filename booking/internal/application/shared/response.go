package shared

type ErrorCodes string

const (
	NotFound            ErrorCodes = "NOT_FOUND"
	CouldNotStoreData   ErrorCodes = "COULD_NOT_STORE_DATA"
	InvalidPersonID     ErrorCodes = "INVALID_PERSON_ID"
	MissingRequiredInfo ErrorCodes = "MISSING_REQUIRED_INFO"
	InvalidEmail        ErrorCodes = "INVALID_EMAIL"
)

type Response struct {
	Success   bool
	Message   string
	ErrorCode ErrorCodes
}
