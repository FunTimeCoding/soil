package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/base"
	"net/http"
	"testing"
)

func TestTheStreamFollowsEveryClaimChange(t *testing.T) {
	s := base.New(t)
	r := openStream(t, s.Port)
	assert.String(t, "[]", readEvent(t, r))
	c := newClient(t, s.Port)
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
		readEvent(t, r),
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
	assert.String(t, "[]", readEvent(t, r))
}

func TestAClaimWithoutAnOwnerIsRejected(t *testing.T) {
	s := base.New(t)
	r, e := newClient(t, s.Port).PutClaimWithResponse(
		context.Background(),
		client.PutClaimJSONRequestBody{Item: "jira-ABC-1"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusBadRequest, r.StatusCode())
}
