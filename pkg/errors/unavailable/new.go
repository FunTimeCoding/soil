package unavailable

import (
	"fmt"
	"time"
)

func New(
	name string,
	since time.Time,
	reason string,
) *UnavailableError {
	return &UnavailableError{
		Message: fmt.Sprintf(
			"%s unreachable since %s: %s",
			name,
			since.Format(time.DateTime),
			reason,
		),
	}
}
