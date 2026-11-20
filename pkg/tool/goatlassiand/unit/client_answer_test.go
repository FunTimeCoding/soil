package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"testing"
)

func TestClientSeesUnknownIssueAsNotFound(t *testing.T) {
	j := mock_client.New(t)
	j.Fail(j.Issue, not_found.Format("Issue does not exist"))
	_, e := client_tester.New(t).Client(j).GetIssue(
		context.Background(),
		"ALFA-1",
		false,
	)
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "Issue does not exist", e.Error())
}

func TestClientSeesRejectedQueryAsValidation(t *testing.T) {
	j := mock_client.New(t)
	j.Fail(j.Search, validation.New("Error in the JQL Query"))
	_, e := client_tester.New(t).Client(j).SearchIssues(
		context.Background(),
		"bravo =",
		nil,
	)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "Error in the JQL Query", e.Error())
}

func TestClientSeesOutageAsNeitherNotFoundNorEmpty(t *testing.T) {
	j := mock_client.New(t)
	j.Fail(j.Search, detail_error.New("Service unavailable", "503"))
	result, e := client_tester.New(t).Client(j).SearchIssues(
		context.Background(),
		"project = CHARLIE",
		nil,
	)
	assert.True(t, e != nil)
	assert.False(t, not_found.Is(e))
	assert.True(t, result == nil)
}

func TestClientReachesGoatlassiand(t *testing.T) {
	c := client_tester.New(t).Client(mock_client.New(t))
	assert.FatalOnError(t, c.Reachable(context.Background()))
}
