package argument_tester

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/identity"
	"testing"
)

func Instance(t *testing.T) *argument.Instance {
	t.Helper()

	return argument.NewInstance(identity.New("gotest", "test tool", "gotest"))
}
