package check_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gocheck"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/option"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
)

func (o *Tester) Run(p *option.Check) *check_result.Result {
	o.t.Helper()
	result, e := gocheck.Run(p)
	assert.FatalOnError(o.t, e)

	return result
}
