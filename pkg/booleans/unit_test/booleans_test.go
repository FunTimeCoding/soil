package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/booleans"
	"testing"
)

func TestToString(t *testing.T) {
	assert.String(t, "1", booleans.ToString(true))
	assert.String(t, "0", booleans.ToString(false))
}
