package unit

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/identity"
	"testing"
)

func testInstance(t *testing.T) *argument.Instance {
	t.Helper()

	return argument.NewInstance(identity.New("gotest", "test tool", "gotest"))
}
