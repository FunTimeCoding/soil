package unit_test

import (
	"github.com/funtimecoding/soil/pkg/git"
	"testing"
)

func TestTags(t *testing.T) {
	git.Tags(git.FindDirectory())
}
