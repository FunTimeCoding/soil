package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/project"
	"path/filepath"
	"testing"
)

func TestTestPath(t *testing.T) {
	assert.String(
		t,
		filepath.Join(git.FindDirectory(), "tmp/test/alfa/bravo"),
		project.TestPath("alfa", "bravo"),
	)
}
