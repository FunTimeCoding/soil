package build

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/git"
	"testing"
)

func TestBuild(t *testing.T) {
	d := git.FindDirectory()

	if d == "" {
		t.Skip("no git directory")
	}

	b := New("a", "b", "c")
	assert.String(t, b.Version, "a")
	assert.String(t, b.GitHash, "b")
	assert.String(t, b.BuildDate, "c")
	assert.True(t, len(Date()) > 0)
	assert.True(t, len(GitHash()) > 0)

	if len(git.Tags(d)) == 0 {
		t.Skip("no git tags")
	}

	assert.True(t, len(GitTag()) > 0)
}
