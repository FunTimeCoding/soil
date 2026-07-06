package run

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "in_progress", InProgress)
	assert.String(t, "queued", Queued)
}
