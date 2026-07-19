package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestCompare(t *testing.T) {
	assertCompare(
		t,
		[]string{},
		[]string{},
		[]string{},
		[]string{},
		[]string{},
	)
	// Add
	assertCompare(
		t,
		[]string{upper.Alfa},
		[]string{},
		[]string{},
		[]string{},
		[]string{upper.Alfa},
	)
	// Remove
	assertCompare(
		t,
		[]string{},
		[]string{upper.Alfa},
		[]string{},
		[]string{upper.Alfa},
		[]string{},
	)
	// Stay
	assertCompare(
		t,
		[]string{},
		[]string{},
		[]string{upper.Alfa},
		[]string{upper.Alfa},
		[]string{upper.Alfa},
	)
}

func assertCompare(
	t *testing.T,
	expectAdd []string,
	expectRemove []string,
	expectStay []string,
	past []string,
	now []string,
) {
	t.Helper()
	add, remove, stay := strings.Compare(past, now)
	assert.Any(t, expectAdd, add)
	assert.Any(t, expectRemove, remove)
	assert.Any(t, expectStay, stay)
}
