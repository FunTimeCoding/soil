package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/linux/systemd/helper"
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	assert.Any(
		t,
		time.Unix(1546300800, 0).UTC(),
		helper.ParseTimestamp("Mon 2019-01-01 00:00:00 UTC"),
	)
}
