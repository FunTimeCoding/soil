package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"testing"
)

func TestKeys(t *testing.T) {
	assert.Strings(
		t,
		[]string{"ABC-123", "XYZ-456"},
		issue.Keys("Message with keys ABC-123 and XYZ-456."),
	)
}
