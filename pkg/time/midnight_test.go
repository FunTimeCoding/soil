package time

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
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
		Midnight(constant.StartOfTime),
	)
}
