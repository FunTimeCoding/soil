package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/failing_jira"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"testing"
)

func TestTheClientSeesAnUnknownIssueAsNotFound(t *testing.T) {
	c := client_tester.New(
		t,
		failing_jira.New(not_found.Format("Issue does not exist")),
	)
	_, e := c.GetIssue(context.Background(), "ALFA-1", false)
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "Issue does not exist", e.Error())
}

func TestTheClientSeesARejectedQueryAsValidation(t *testing.T) {
	c := client_tester.New(
		t,
		failing_jira.New(validation.New("Error in the JQL Query")),
	)
	_, e := c.SearchIssues(context.Background(), "bravo =", nil)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "Error in the JQL Query", e.Error())
}

func TestTheClientSeesAnOutageAsNeitherNotFoundNorEmpty(t *testing.T) {
	c := client_tester.New(
		t,
		failing_jira.New(detail_error.New("Service unavailable", "503")),
	)
	result, e := c.SearchIssues(context.Background(), "project = CHARLIE", nil)
	assert.True(t, e != nil)
	assert.False(t, not_found.Is(e))
	assert.True(t, result == nil)
}

func TestTheClientReachesGoatlassiand(t *testing.T) {
	c := client_tester.New(t, mock_client.New())
	assert.FatalOnError(t, c.Reachable(context.Background()))
}
