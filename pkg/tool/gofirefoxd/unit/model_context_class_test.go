package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	firefoxConstant "github.com/funtimecoding/soil/pkg/firefox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/model_context_tester"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/tool_case"
	"testing"
)

func TestEveryToolCoversEveryTool(t *testing.T) {
	assert.Count(
		t,
		len(tool_case.Tools()),
		model_context_tester.New(t).Client.ListTools(),
	)
}

func TestMissingExtensionIsAnsweredNotReported(t *testing.T) {
	for _, c := range tool_case.Tools() {
		o := model_context_tester.New(t)
		o.MockClient.Fail(
			c.Method(o.MockClient),
			unavailable.Format(firefoxConstant.NotConnected),
		)
		assert.StringContains(
			t,
			"extension not connected",
			o.Client.MustCallToolError(c.Name, c.Arguments),
		)
		assert.Count(t, 0, o.Reporter.Events())
	}
}

func TestRefusalIsAnsweredNotReported(t *testing.T) {
	for _, c := range tool_case.Tools() {
		o := model_context_tester.New(t)
		o.MockClient.Fail(
			c.Method(o.MockClient),
			not_found.Format("%s", "selector matched nothing: #go"),
		)
		assert.StringContains(
			t,
			"matched nothing",
			o.Client.MustCallToolError(c.Name, c.Arguments),
		)
		assert.Count(t, 0, o.Reporter.Events())
	}
}

func TestUnexpectedExtensionErrorIsReported(t *testing.T) {
	for _, c := range tool_case.Tools() {
		o := model_context_tester.New(t)
		o.MockClient.Fail(
			c.Method(o.MockClient),
			unexpected.Format("extension: %s", "Invalid tab ID: 1"),
		)
		o.Client.MustCallToolError(c.Name, c.Arguments)
		assert.Count(t, 1, o.Reporter.Events())
	}
}
