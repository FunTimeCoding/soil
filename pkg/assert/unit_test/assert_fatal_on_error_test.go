package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFatalOnError(t *testing.T) {
	assert.FatalOnError(t, nil)
}
