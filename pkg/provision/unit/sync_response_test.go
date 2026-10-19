package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/provision/model_context"
	"github.com/funtimecoding/soil/pkg/provision/types/update"
	"github.com/mark3labs/mcp-go/mcp"
	"testing"
)

func TestAQueuedSyncIsAnsweredQuietly(t *testing.T) {
	m := memory.New()
	r, e := model_context.SyncResponse(
		m,
		nil,
		conflict.Format("sync already queued"),
	)
	assert.Nil(t, e)
	assert.String(t, "sync already queued", r.Content[0].(mcp.TextContent).Text)
	assert.Count(t, 0, m.Events())
}

func TestAnUnexpectedSyncErrorIsCaptured(t *testing.T) {
	m := memory.New()
	_, e := model_context.SyncResponse(m, nil, errors.New("channel closed"))
	assert.Nil(t, e)
	assert.Count(t, 1, m.Events())
}

func TestAFailedSyncIsNotCapturedTwice(t *testing.T) {
	m := memory.New()
	result := update.NewResult()
	result.Error = errors.New("sync failed")
	r, e := model_context.SyncResponse(m, result, nil)
	assert.Nil(t, e)
	assert.String(t, "sync failed", r.Content[0].(mcp.TextContent).Text)
	assert.Count(t, 0, m.Events())
}
