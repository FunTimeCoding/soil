package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/model_context_tester"
	"testing"
)

func TestAQuietClassIsAnsweredWithItsOwnText(t *testing.T) {
	s := model_context_tester.New(t)
	assert.String(
		t,
		"scope name is reserved: all",
		s.MustCallToolError(
			constant.Profile,
			map[string]any{constant.Scope: constant.AllScope},
		),
	)
}
