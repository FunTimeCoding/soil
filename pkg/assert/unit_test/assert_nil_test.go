package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestNil(t *testing.T) {
	assert.Nil(t, nil)
}
