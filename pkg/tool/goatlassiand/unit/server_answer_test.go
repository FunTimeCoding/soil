package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/failing_jira"
	"testing"
)

func TestAnUnknownIssueIsAnsweredNotFoundUncaptured(t *testing.T) {
	m, e := getIssue(t, not_found.Format("Issue does not exist"))
	assert.True(t, not_found.Is(e))
	assert.Count(t, 0, m.Events())
}

func TestARejectedQueryIsAnsweredBadRequestUncaptured(t *testing.T) {
	c, m := client_tester.NewReported(
		t,
		failing_jira.New(validation.New("Error in the JQL Query")),
	)
	_, e := c.SearchIssues(context.Background(), "bravo =", nil)
	assert.True(t, validation.Is(e))
	assert.Count(t, 0, m.Events())
}
