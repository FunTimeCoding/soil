package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestLess(t *testing.T) {
	assert.Less(t, 1, 0)
	// Unhappy more
	t1 := &testing.T{}
	assert.Less(t1, 0, 1)

	if !t1.Failed() {
		fmt.Println("unhappy more")
		t.Fail()
	}

	// Unhappy equal
	t2 := &testing.T{}
	assert.Less(t2, 1, 1)

	if !t2.Failed() {
		fmt.Println("unhappy equal")
		t.Fail()
	}
}
