package toolchain

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	"testing"
)

func cachedRun(
	t *testing.T,
	root string,
	index string,
	fresh bool,
) (*check_result.Result, string) {
	t.Helper()
	o := option.New()
	o.Root = root
	o.Index = index
	o.Fresh = fresh
	var result *check_result.Result
	printed := testutil.CaptureStdout(
		func() {
			result = checked(t, o)
			gocheck.Print(result)
		},
	)

	return result, printed
}
