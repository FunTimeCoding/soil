package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestRound(t *testing.T) {
	assert.Round(t, 1.1, 1.14, 1)
	assert.Round(t, 1.15, 1.154, 2)
}
