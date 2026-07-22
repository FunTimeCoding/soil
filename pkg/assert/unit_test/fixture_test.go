package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"testing"
)

func TestFixture(t *testing.T) {
	assert.Suffix(
		t,
		"soil/fixture/example.txt",
		fixture.Path("example.txt"),
	)
}

func TestFileExists(t *testing.T) {
	assert.FileExists(t, "./fixture_test.go")
}
