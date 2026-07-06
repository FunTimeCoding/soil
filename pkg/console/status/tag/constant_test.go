package tag

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "age", Age)
	assert.String(t, "assignee", Assignee)
	assert.String(t, "concerns", Concerns)
	assert.String(t, "description", Description)
	assert.String(t, "graph", Graph)
	assert.String(t, "identifier", Identifier)
	assert.String(t, "key", Key)
	assert.String(t, "labels", Labels)
	assert.String(t, "markdown", Markdown)
	assert.String(t, "runbook", Runbook)
	assert.String(t, "score", Score)
	assert.String(t, "status", Status)
	assert.String(t, "timestamp", Timestamp)
	assert.String(t, "type", Type)
	assert.String(t, "wiki", Wiki)
}
