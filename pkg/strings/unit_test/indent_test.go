package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/indent"
	"testing"
)

func TestF(t *testing.T) {
	assert.NotNil(t, indent.F)
}
