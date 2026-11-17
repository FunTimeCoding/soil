package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
)

func TestQuietClassIsAnsweredWithItsOwnText(t *testing.T) {
	assert.String(
		t,
		"scope name is reserved: all",
		base.New(t).Tools().MustCallToolError(
			constant.Profile,
			map[string]any{constant.Scope: constant.AllScope},
		),
	)
}
