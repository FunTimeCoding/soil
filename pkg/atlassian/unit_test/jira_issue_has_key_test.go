package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"testing"
)

func TestHasKey(t *testing.T) {
	assert.True(
		t,
		issue.HasKey("Message with key ABC-123."),
	)
}
