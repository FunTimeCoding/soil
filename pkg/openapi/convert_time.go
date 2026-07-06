package openapi

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-openapi/strfmt"
	"time"
)

func ConvertTime(t *strfmt.DateTime) *time.Time {
	result, e := time.Parse(time.RFC3339, t.String())
	errors.PanicOnError(e)

	return &result
}
