package alert

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestAlert(t *testing.T) {
	assert.Count(t, 2, Statuses)
	assert.Count(t, 5, Priorities)
	assert.Count(t, 1, SkipDetail)
	assert.Count(t, 0, CondenseSkipFields)
}
