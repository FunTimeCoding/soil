package remediable

import "fmt"

func Wrap(
	e error,
	remedy string,
	sentence string,
) *RemediableError {
	return &RemediableError{
		Name:    remedy,
		Message: fmt.Sprintf("%s: %v", sentence, e),
		cause:   e,
	}
}
