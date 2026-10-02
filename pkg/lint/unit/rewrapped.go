package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/reflow"
	"testing"
)

func rewrapped(
	t *testing.T,
	content string,
) string {
	result, e := reflow.Reflow(content, constant.FixtureReflowWidth)
	assert.FatalOnError(t, e)

	return result
}
