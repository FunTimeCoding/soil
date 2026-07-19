package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/openapi"
	"github.com/go-openapi/strfmt"
	"testing"
	"time"
)

func TestConvertTime(t *testing.T) {
	i := strfmt.NewDateTime()
	errors.PanicOnError(i.Scan(constant.StartOfTime))
	assert.String(
		t,
		"1970-01-01T00:00:00Z",
		openapi.ConvertTime(&i).Format(time.RFC3339),
	)
}
