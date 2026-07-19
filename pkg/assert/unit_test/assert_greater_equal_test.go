package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestGreaterEqual(t *testing.T) {
	assert.GreaterEqual(t, 0, 1)
	assert.GreaterEqual(t, 1, 1)
	// Unhappy less
	t1 := &testing.T{}
	assert.GreaterEqual(t1, 1, 0)

	if !t1.Failed() {
		fmt.Println("unhappy less")
		t.Fail()
	}
}
