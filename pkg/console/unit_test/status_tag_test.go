package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"testing"
)

func TestStatusTagConstant(t *testing.T) {
	assert.String(t, "age", tag.Age)
	assert.String(t, "assignee", tag.Assignee)
	assert.String(t, "concerns", tag.Concerns)
	assert.String(t, "description", tag.Description)
	assert.String(t, "graph", tag.Graph)
	assert.String(t, "identifier", tag.Identifier)
	assert.String(t, "key", tag.Key)
	assert.String(t, "labels", tag.Labels)
	assert.String(t, "markdown", tag.Markdown)
	assert.String(t, "runbook", tag.Runbook)
	assert.String(t, "score", tag.Score)
	assert.String(t, "status", tag.Status)
	assert.String(t, "timestamp", tag.Timestamp)
	assert.String(t, "type", tag.Type)
	assert.String(t, "wiki", tag.Wiki)
}
