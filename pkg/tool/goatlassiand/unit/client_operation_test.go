package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/web/detail_error"
	"testing"
)

func TestClientReadsBody(t *testing.T) {
	spaces, e := client_tester.New(t).Client(mock_client.New(t)).ListSpaces(
		context.Background(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, spaces)
}

func TestClientAcceptsNoContent(t *testing.T) {
	assert.FatalOnError(
		t,
		client_tester.New(t).Client(mock_client.New(t)).TransitionIssue(
			context.Background(),
			"ALFA-1",
			"bravo",
		),
	)
}

func TestClientRefusesFailedWrite(t *testing.T) {
	j := mock_client.New(t)
	j.Fail(j.Transition, detail_error.New("Transition is not valid", "400"))
	e := client_tester.New(t).Client(j).TransitionIssue(
		context.Background(),
		"ALFA-1",
		"bravo",
	)
	assert.True(t, e != nil)
	assert.StringContains(t, "Transition is not valid", e.Error())
}
