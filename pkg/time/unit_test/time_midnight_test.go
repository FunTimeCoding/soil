package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	library "github.com/funtimecoding/soil/pkg/time"
	"testing"
	"time"
)

func TestMidnight(t *testing.T) {
	assert.Any(
		t,
		time.Date(
			1970,
			1,
			1,
			0,
			0,
			0,
			0,
			time.Local,
		),
		library.Midnight(constant.StartOfTime),
	)
}
