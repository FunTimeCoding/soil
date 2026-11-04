package reported

import "fmt"

func New(
	operation string,
	reason string,
	event string,
) *ReportedError {
	return &ReportedError{
		Message: fmt.Sprintf(
			"%s: %s (reported as %s)",
			operation,
			reason,
			event,
		),
		Event: event,
	}
}
