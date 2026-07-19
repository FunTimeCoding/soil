package unit_test

import (
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestExitOnEmpty(t *testing.T) {
	system.ExitOnEmpty(1, "notEmpty", "This should not exit")
}
