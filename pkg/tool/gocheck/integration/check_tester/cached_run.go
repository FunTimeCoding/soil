package check_tester

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
)

func (o *Tester) CachedRun(
	root string,
	index string,
	fresh bool,
) (*check_result.Result, string) {
	o.t.Helper()
	p := option.New()
	p.Root = root
	p.Index = index
	p.Fresh = fresh
	var result *check_result.Result
	printed := testutil.CaptureStdout(
		func() {
			result = o.Run(p)
			gocheck.Print(result)
		},
	)

	return result, printed
}
