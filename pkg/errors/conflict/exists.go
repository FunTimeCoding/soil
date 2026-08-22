package conflict

import "fmt"

func Exists(
	kind string,
	identifier any,
) *ConflictError {
	return &ConflictError{
		Message: fmt.Sprintf("%s already exists: %v", kind, identifier),
	}
}
