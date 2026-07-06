package time

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-openapi/strfmt"
	"time"
)

func Scan(t time.Time) strfmt.DateTime {
	result := strfmt.NewDateTime()
	errors.PanicOnError(result.Scan(t))

	return result
}
