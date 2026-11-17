package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/base"
	"testing"
)

func TestUnknownEntryIsAnsweredQuietlyOverModelContext(t *testing.T) {
	s := base.New(t)
	c := s.Tools()
	assert.String(
		t,
		"entry not found: 9999",
		c.MustCallToolError(
			constant.DeleteEntry,
			map[string]any{constant.Identifier: float64(9999)},
		),
	)
	assert.String(
		t,
		"entry not found: 9999",
		c.MustCallToolError(
			constant.UpdateEntry,
			map[string]any{
				constant.Identifier: float64(9999),
				constant.Action:     "x",
			},
		),
	)
	assert.Integer(t, 0, s.Captured())
}
