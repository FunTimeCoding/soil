package lint_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/reflow"
	"testing"
)

func Rewrapped(
	t *testing.T,
	content string,
) string {
	t.Helper()
	result, e := reflow.Reflow(content, constant.FixtureReflowWidth)
	assert.FatalOnError(t, e)

	return result
}
