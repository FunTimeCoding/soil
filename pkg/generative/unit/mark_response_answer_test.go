package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
	"testing"
)

func TestAnswerFailsAnUncapturedClassQuietly(t *testing.T) {
	m := memory.New()
	r, e := response.Answer(m, not_found.Format("post not found: 8"))
	assert.Nil(t, e)
	assert.True(t, r.IsError)
	assert.String(t, "post not found: 8", r.Content[0].(mcp.TextContent).Text)
	assert.Count(t, 0, m.Events())
}

func TestAnswerCapturesACapturedClassWithItsMessage(t *testing.T) {
	m := memory.New()
	r, e := response.Answer(m, not_configured.Format("no instance configured"))
	assert.Nil(t, e)
	assert.True(t, r.IsError)
	assert.StringContains(
		t,
		"no instance configured",
		r.Content[0].(mcp.TextContent).Text,
	)
	assert.Count(t, 1, m.Events())
}

func TestAnswerCapturesTheUnclassifiedWithItsText(t *testing.T) {
	m := memory.New()
	r, e := response.Answer(m, errors.New("disk full"))
	assert.Nil(t, e)
	assert.StringContains(
		t,
		`"error": "disk full"`,
		r.Content[0].(mcp.TextContent).Text,
	)
	assert.Count(t, 1, m.Events())
}

func TestAnswerWithPutsTheContextBeforeAQuietClass(t *testing.T) {
	m := memory.New()
	r, e := response.AnswerWith(
		m,
		not_found.Format("post not found: 8"),
		"load post",
	)
	assert.Nil(t, e)
	assert.String(
		t,
		"load post: post not found: 8",
		r.Content[0].(mcp.TextContent).Text,
	)
	assert.Count(t, 0, m.Events())
}

func TestAnswerWithPutsTheContextBeforeAFault(t *testing.T) {
	m := memory.New()
	r, e := response.AnswerWith(m, errors.New("disk full"), "load post")
	assert.Nil(t, e)
	assert.StringContains(
		t,
		`"error": "load post: disk full"`,
		r.Content[0].(mcp.TextContent).Text,
	)
	assert.Count(t, 1, m.Events())
}
