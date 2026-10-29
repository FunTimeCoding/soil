package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
	"github.com/mark3labs/mcp-go/mcp"
	"testing"
)

func TestAnUnknownLinkIsAnsweredQuietlyOverModelContext(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = not_found.Format("resource not found: links/%d", 12)
	m := memory.New()
	s := model_context.New(
		service_tester.Service(t, c, m),
		m,
		mock_recorder.New(),
	)
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
