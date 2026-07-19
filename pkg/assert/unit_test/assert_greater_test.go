package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestGreater(t *testing.T) {
	assert.Greater(t, 0, 1)
	// Unhappy less
	t1 := &testing.T{}
	assert.Greater(t1, 1, 0)

	if !t1.Failed() {
		fmt.Println("unhappy less")
		t.Fail()
	}

	// Unhappy equal
	t2 := &testing.T{}
	assert.Greater(t2, 1, 1)

	if !t2.Failed() {
		fmt.Println("unhappy equal")
		t.Fail()
	}
}
