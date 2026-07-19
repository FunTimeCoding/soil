package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

type Fixture struct {
	Value string
}

func TestAny(t *testing.T) {
	assert.Any(t, &Fixture{Value: "a"}, &Fixture{Value: "a"})
}
