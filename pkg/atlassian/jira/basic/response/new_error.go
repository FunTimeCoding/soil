package response

func NewError(
	errorMessages []string,
	errors map[string]string,
) *Error {
	return &Error{ErrorMessages: errorMessages, Errors: errors}
}
