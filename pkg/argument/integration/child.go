package integration

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"os"
	"testing"
)

func child(
	t *testing.T,
	arguments string,
) *run.Run {
	t.Helper()
	result := run.New().NoPanic()
	result.Environment(constant.FixtureArgumentsEnvironment, arguments)
	result.Start(os.Args[0], fmt.Sprintf("-test.run=^%s$", t.Name()))

	return result
}
