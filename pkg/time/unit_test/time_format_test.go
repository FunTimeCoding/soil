package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/time"
	"testing"
)

func TestFormat(t *testing.T) {
	assert.String(
		t,
		"1970-01-01 00:00",
		time.Format(constant.StartOfTime),
	)
}
