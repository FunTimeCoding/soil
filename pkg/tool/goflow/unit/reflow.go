package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goflow"
	"github.com/funtimecoding/soil/pkg/tool/goflow/constant"
	"testing"
)

func reflow(
	t *testing.T,
	content string,
) string {
	result, e := goflow.Reflow(content, constant.FixtureWidth)
	assert.FatalOnError(t, e)

	return result
}
