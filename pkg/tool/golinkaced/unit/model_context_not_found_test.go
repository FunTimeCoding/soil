package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/missing_link"
	"github.com/mark3labs/mcp-go/mcp"
	"testing"
)

func TestAnUnknownLinkIsAnsweredQuietlyOverModelContext(t *testing.T) {
	m := memory.New()
	s := model_context.New(missing_link.New(), nil, m, mock_recorder.New())
	r, e := s.Delete(
		context.Background(),
		mcp.CallToolRequest{},
		argument.Delete{Type: "link", Identifier: 12},
	)
	assert.FatalOnError(t, e)
	assert.True(t, r.IsError)
	assert.String(
		t,
		"resource not found: links/12",
		r.Content[0].(mcp.TextContent).Text,
	)
	assert.Count(t, 0, m.Events())
}
