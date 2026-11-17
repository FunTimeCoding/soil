package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	firefoxConstant "github.com/funtimecoding/soil/pkg/firefox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/tool_case"
	"testing"
)

func TestEveryToolCoversEveryTool(t *testing.T) {
	assert.Count(t, len(tool_case.Tools()), base.New(t).Tools().ListTools())
}

func TestMissingExtensionIsAnsweredNotReported(t *testing.T) {
	for _, c := range tool_case.Tools() {
		s := base.New(t)
		s.Source.Fail(
			c.Method(s.Source),
			unavailable.Format(firefoxConstant.NotConnected),
		)
		assert.StringContains(
			t,
			"extension not connected",
			s.Tools().MustCallToolError(c.Name, c.Arguments),
		)
		assert.Count(t, 0, s.Reporter.Events())
	}
}

func TestRefusalIsAnsweredNotReported(t *testing.T) {
	for _, c := range tool_case.Tools() {
		s := base.New(t)
		s.Source.Fail(
			c.Method(s.Source),
			not_found.Format("%s", "selector matched nothing: #go"),
		)
		assert.StringContains(
			t,
			"matched nothing",
			s.Tools().MustCallToolError(c.Name, c.Arguments),
		)
		assert.Count(t, 0, s.Reporter.Events())
	}
}

func TestUnexpectedExtensionErrorIsReported(t *testing.T) {
	for _, c := range tool_case.Tools() {
		s := base.New(t)
		s.Source.Fail(
			c.Method(s.Source),
			unexpected.Format("extension: %s", "Invalid tab ID: 1"),
		)
		s.Tools().MustCallToolError(c.Name, c.Arguments)
		assert.Count(t, 1, s.Reporter.Events())
	}
}
