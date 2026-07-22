package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestDeviate(t *testing.T) {
	assert.Deviate(t, 1, 1.1, 0.1)
	assert.Deviate(t, -1, -1.1, 0.1)
}

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

func TestRound(t *testing.T) {
	assert.Round(t, 1.1, 1.14, 1)
	assert.Round(t, 1.15, 1.154, 2)
}
