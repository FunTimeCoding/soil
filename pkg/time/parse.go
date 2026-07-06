package time

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func Parse(
	layout string,
	s string,
) time.Time {
	result, e := time.Parse(layout, s)
	errors.PanicOnError(e)

	return result
}
