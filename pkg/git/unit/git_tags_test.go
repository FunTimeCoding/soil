package unit

import (
	"github.com/funtimecoding/soil/pkg/git"
	"testing"
)

func TestTags(t *testing.T) {
	git.Tags(git.FindDirectory())
}
