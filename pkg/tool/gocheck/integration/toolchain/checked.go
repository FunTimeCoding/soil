package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	"testing"
)

func checked(
	t *testing.T,
	o *option.Check,
) *check_result.Result {
	t.Helper()
	result, e := gocheck.Run(o)
	assert.FatalOnError(t, e)

	return result
}
