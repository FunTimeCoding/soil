package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/minute"
	"testing"
)

func TestMinuteReadable(t *testing.T) {
	assert.String(t, "1 minute", minute.Readable(1))
	assert.String(t, "1 hour", minute.Readable(60))
}
