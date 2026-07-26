package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"testing"
)

func TestStatusTagConstant(t *testing.T) {
	assert.String(t, "age", constant.TagAge)
	assert.String(t, "assignee", constant.TagAssignee)
	assert.String(t, "concerns", constant.TagConcerns)
	assert.String(t, "description", constant.TagDescription)
	assert.String(t, "graph", constant.TagGraph)
	assert.String(t, "identifier", constant.TagIdentifier)
	assert.String(t, "key", constant.TagKey)
	assert.String(t, "labels", constant.TagLabels)
	assert.String(t, "markdown", constant.TagMarkdown)
	assert.String(t, "runbook", constant.TagRunbook)
	assert.String(t, "score", constant.TagScore)
	assert.String(t, "status", constant.TagStatus)
	assert.String(t, "timestamp", constant.TagTimestamp)
	assert.String(t, "type", constant.TagType)
	assert.String(t, "wiki", constant.TagWiki)
}
