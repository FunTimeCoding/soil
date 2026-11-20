package scope_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goanalyze/constant"
	"os"
	"strconv"
)

func (o *Tester) Child(
	directory string,
	full bool,
) *run.Run {
	o.t.Helper()
	result := run.New().NoPanic()
	result.Environment(constant.FixtureDirectoryEnvironment, directory)
	result.Environment(
		constant.FixtureFullEnvironment,
		strconv.FormatBool(full),
	)
	result.Start(os.Args[0], fmt.Sprintf("-test.run=^%s$", o.t.Name()))

	return result
}
