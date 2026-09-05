package unit

import (
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestSystemRun(t *testing.T) {
	system.Run("ls")
}
