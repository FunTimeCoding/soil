package unit

import (
	"github.com/funtimecoding/soil/pkg/git"
	"testing"
)

func TestStatus(t *testing.T) {
	git.Status(git.FindDirectory())
}
