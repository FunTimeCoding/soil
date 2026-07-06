package floats

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestToString(t *testing.T) {
	assert.String(t, "0", ToString(0))
	assert.String(t, "0.1", ToString(0.1))
	assert.String(t, "1", ToString(1))
	assert.String(t, "1.1", ToString(1.1))
	assert.String(t, "1.12", ToString(1.12))
}

func TestToStringRounded(t *testing.T) {
	assert.String(t, "0.1", ToStringRounded(0.11))
	assert.String(t, "0.2", ToStringRounded(0.19))
}
