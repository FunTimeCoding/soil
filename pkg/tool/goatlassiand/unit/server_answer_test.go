package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/client_tester"
	"testing"
)

func TestUnknownIssueIsAnsweredNotFoundUncaptured(t *testing.T) {
	m, e := client_tester.New(t).GetIssue(
		not_found.Format("Issue does not exist"),
	)
	assert.True(t, not_found.Is(e))
	assert.Count(t, 0, m.Events())
}

func TestRejectedQueryIsAnsweredBadRequestUncaptured(t *testing.T) {
	j := mock_client.New(t)
	j.Fail(j.Search, validation.New("Error in the JQL Query"))
	c, m := client_tester.New(t).Reported(j)
	_, e := c.SearchIssues(context.Background(), "bravo =", nil)
	assert.True(t, validation.Is(e))
	assert.Count(t, 0, m.Events())
}
