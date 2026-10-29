package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/failing_jira"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"testing"
)

func TestTheClientReadsABody(t *testing.T) {
	spaces, e := client_tester.New(t, mock_client.New()).ListSpaces(
		context.Background(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, spaces)
}

func TestTheClientAcceptsNoContent(t *testing.T) {
	assert.FatalOnError(
		t,
		client_tester.New(t, mock_client.New()).TransitionIssue(
			context.Background(),
			"ALFA-1",
			"bravo",
		),
	)
}

func TestTheClientRefusesAFailedWrite(t *testing.T) {
	e := client_tester.New(
		t,
		failing_jira.New(detail_error.New("Transition is not valid", "400")),
	).TransitionIssue(context.Background(), "ALFA-1", "bravo")
	assert.True(t, e != nil)
	assert.StringContains(t, "Transition is not valid", e.Error())
}
