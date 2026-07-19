package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/dash"
	"testing"
)

func TestToUnderscore(t *testing.T) {
	assert.String(t, "a_b_c", dash.ToUnderscore("a-b-c"))
}
