package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestLessEqual(t *testing.T) {
	assert.LessEqual(t, 1, 0)
	assert.LessEqual(t, 1, 1)
	// Unhappy more
	t1 := &testing.T{}
	assert.LessEqual(t1, 0, 1)

	if !t1.Failed() {
		fmt.Println("unhappy more")
		t.Fail()
	}
}
