package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/service_tester"
	"github.com/mark3labs/mcp-go/mcp"
	"testing"
)

func TestAnUnknownIssueIsAnsweredQuietly(t *testing.T) {
	m := memory.New()
	s := model_context.New(
		service_tester.Service(mock_client.New()),
		"sentry.example",
		m,
		mock_recorder.New(),
	)
	r, e := s.GetIssue(
		context.Background(),
		mcp.CallToolRequest{},
		argument.GetIssue{Identifier: "100"},
	)
	assert.FatalOnError(t, e)
	assert.True(t, r.IsError)
	assert.String(
		t,
		"issue not found: 100",
		r.Content[0].(mcp.TextContent).Text,
	)
	assert.Count(t, 0, m.Events())
}
