package unit

import (
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestEnsureFileDeleted(t *testing.T) {
	system.EnsureFileDeleted("neverExistingFile")
}
