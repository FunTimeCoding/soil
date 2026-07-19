package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/floats"
	"testing"
)

func TestToString(t *testing.T) {
	assert.String(t, "0", floats.ToString(0))
	assert.String(t, "0.1", floats.ToString(0.1))
	assert.String(t, "1", floats.ToString(1))
	assert.String(t, "1.1", floats.ToString(1.1))
	assert.String(t, "1.12", floats.ToString(1.12))
}

func TestToStringRounded(t *testing.T) {
	assert.String(t, "0.1", floats.ToStringRounded(0.11))
	assert.String(t, "0.2", floats.ToStringRounded(0.19))
}
