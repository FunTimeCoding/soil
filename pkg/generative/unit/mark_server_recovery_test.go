package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/unit/identified_reporter"
	"github.com/funtimecoding/soil/pkg/generative/unit/server_tester"
	"github.com/mark3labs/mcp-go/mcp"
	"testing"
)

func TestPanickingToolAnswersReadableErrorWithEventIdentifier(t *testing.T) {
	r := identified_reporter.New("echo")
	c := server_tester.Panicking(t, r)
	q := mcp.CallToolRequest{}
	q.Params.Name = "golf"
	result, e := c.CallTool(t.Context(), q)
	assert.FatalOnError(t, e)
	assert.True(t, result.IsError)
	assert.String(
		t,
		"{\n\t\"error\": \"golf panicked\",\n\t\"event_identifier\": \"echo\"\n}",
		result.Content[0].(mcp.TextContent).Text,
	)
	assert.Count(t, 1, r.Events())
}

func TestPanickingResourceAnswersErrorWithEventIdentifier(t *testing.T) {
	r := identified_reporter.New("foxtrot")
	c := server_tester.Panicking(t, r)
	q := mcp.ReadResourceRequest{}
	q.Params.URI = "test://charlie"
	_, e := c.ReadResource(t.Context(), q)
	assert.StringContains(
		t,
		"test://charlie panicked, event foxtrot",
		e.Error(),
	)
	assert.Count(t, 1, r.Events())
}
