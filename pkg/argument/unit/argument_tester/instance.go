package argument_tester

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/identity"
)

func (o *Tester) Instance() *argument.Instance {
	o.t.Helper()

	return argument.NewInstance(identity.New("gotest", "test tool", "gotest"))
}
