package argument_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"os"
)

func (o *Tester) Child(arguments string) *run.Run {
	o.t.Helper()
	result := run.New().NoPanic()
	result.Environment(constant.FixtureArgumentsEnvironment, arguments)
	result.Start(os.Args[0], fmt.Sprintf("-test.run=^%s$", o.t.Name()))

	return result
}
