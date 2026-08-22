package timeout

import (
	"context"
	"errors"
)

func Is(e error) bool {
	var target *TimeoutError

	return errors.As(e, &target) || errors.Is(e, context.DeadlineExceeded)
}
