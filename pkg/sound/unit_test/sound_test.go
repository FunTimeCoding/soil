package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/sound"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "afplay", sound.Afplay)
}
