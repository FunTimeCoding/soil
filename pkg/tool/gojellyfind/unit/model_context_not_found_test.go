package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"testing"
)

func TestAnUnknownItemIsAnsweredQuietlyOverModelContext(t *testing.T) {
	m := memory.New()
	s := model_context.New(mock_client.New(), m, mock_recorder.New())
	r, e := s.GetItem(
		context.Background(),
		mcp.CallToolRequest{},
		argument.GetItem{Identifier: "charlie"},
	)
	assert.FatalOnError(t, e)
	assert.True(t, r.IsError)
	assert.String(
		t,
		"item not found: charlie",
		r.Content[0].(mcp.TextContent).Text,
	)
	assert.Count(t, 0, m.Events())
}
