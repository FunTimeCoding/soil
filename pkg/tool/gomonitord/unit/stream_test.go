package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/stream_tester"
	"net/http"
	"testing"
)

func TestStreamFollowsEveryClaimChange(t *testing.T) {
	s := base.New(t)
	r := stream_tester.Open(t, s.Served())
	assert.String(t, "[]", stream_tester.ReadEvent(t, r))
	c := s.GeneratedClient()
	claimed, e := c.PutClaimWithResponse(
		context.Background(),
		client.PutClaimJSONRequestBody{
			Item:  "jira-ABC-1",
			Owner: "alfa@host.example",
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, claimed.StatusCode())
	assert.StringContains(
		t,
		`"item":"jira-ABC-1","owner":"alfa@host.example"`,
		stream_tester.ReadEvent(t, r),
	)
	released, f := c.DeleteClaimWithResponse(
		context.Background(),
		&client.DeleteClaimParams{
			Item:  "jira-ABC-1",
			Owner: "alfa@host.example",
		},
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, http.StatusNoContent, released.StatusCode())
	assert.String(t, "[]", stream_tester.ReadEvent(t, r))
}

func TestClaimWithoutOwnerIsRejected(t *testing.T) {
	s := base.New(t)
	r, e := s.GeneratedClient().PutClaimWithResponse(
		context.Background(),
		client.PutClaimJSONRequestBody{Item: "jira-ABC-1"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusBadRequest, r.StatusCode())
}
