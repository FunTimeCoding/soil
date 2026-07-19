package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestNotNil(t *testing.T) {
	assert.NotNil(t, 1)
}
