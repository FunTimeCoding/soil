package helper

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	assert.Any(
		t,
		time.Unix(1546300800, 0).UTC(),
		ParseTimestamp("Mon 2019-01-01 00:00:00 UTC"),
	)
}
