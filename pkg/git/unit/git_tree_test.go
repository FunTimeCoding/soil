package unit

import (
	"github.com/funtimecoding/soil/pkg/git"
	"testing"
)

func TestTree(t *testing.T) {
	git.Tree(git.FindDirectory())
}
