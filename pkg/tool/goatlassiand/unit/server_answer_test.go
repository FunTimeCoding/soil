package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	generated "github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/failing_jira"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"testing"
)

func TestAnUnknownIssueIsAnsweredNotFound(t *testing.T) {
	m := memory.New()
	r, e := server.New(
		service_tester.NewService(
			t,
			failing_jira.New(not_found.Format("Issue does not exist")),
			mock_client.New(),
			nil,
		),
		m,
	).GetIssue(
		context.Background(),
		generated.GetIssueRequestObject{Key: "ALFA-1"},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.GetIssue404JSONResponse{Error: "Issue does not exist"},
		r,
	)
	assert.Count(t, 0, m.Events())
}

func TestARejectedQueryIsAnsweredBadRequest(t *testing.T) {
	m := memory.New()
	r, e := server.New(
		service_tester.NewService(
			t,
			failing_jira.New(validation.New("Error in the JQL Query")),
			mock_client.New(),
			nil,
		),
		m,
	).SearchIssues(
		context.Background(),
		generated.SearchIssuesRequestObject{
			Params: generated.SearchIssuesParams{Query: "bravo ="},
		},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.SearchIssues400JSONResponse{Error: "Error in the JQL Query"},
		r,
	)
	assert.Count(t, 0, m.Events())
}
