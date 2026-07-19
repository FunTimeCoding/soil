package unit_test

import (
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestEnsureFileDeleted(t *testing.T) {
	system.EnsureFileDeleted("neverExistingFile")
}
