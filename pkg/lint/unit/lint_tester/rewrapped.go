package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/reflow"
)

func (o *Tester) Rewrapped(content string) string {
	o.t.Helper()
	result, e := reflow.Reflow(content, constant.FixtureReflowWidth)
	assert.FatalOnError(o.t, e)

	return result
}
