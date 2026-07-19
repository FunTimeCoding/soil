package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/run"
	"testing"
)

func TestRunConstant(t *testing.T) {
	assert.String(t, "in_progress", run.InProgress)
	assert.String(t, "queued", run.Queued)
}
