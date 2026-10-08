package web_service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/integration/web_service_tester"
	"net/http"
	"testing"
)

func TestFilterWebService(t *testing.T) {
	o := web_service_tester.New(t)
	c := o.Client
	x := context.Background()
	system1 := "worker1"
	system2 := "worker2"
	service1 := "nginx"
	service2 := "postgres"
	user1 := "alice"
	user2 := "bob"
	desc := "test"

	for _, b := range []client.PostEntryJSONRequestBody{
		{
			Action:      "restart",
			User:        user1,
			System:      &system1,
			Service:     &service1,
			Description: &desc,
		},
		{
			Action:      "backup",
			User:        user2,
			System:      &system2,
			Service:     &service2,
			Description: &desc,
		},
		{
			Action:      "deploy",
			User:        user1,
			System:      &system2,
			Service:     &service1,
			Description: &desc,
		},
	} {
		_, e := c.PostEntryWithResponse(x, b)
		assert.FatalOnError(t, e)
	}

	all, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{})
	assert.FatalOnError(t, e)
	assert.Count(t, 3, *all.JSON200)
	bySys1, e := c.GetEntriesWithResponse(
		x,
		&client.GetEntriesParams{System: &system1},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *bySys1.JSON200)
	assert.String(t, "restart", (*bySys1.JSON200)[0].Action)
	bySys2, e := c.GetEntriesWithResponse(
		x,
		&client.GetEntriesParams{System: &system2},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, *bySys2.JSON200)
	byAlice, e := c.GetEntriesWithResponse(
		x,
		&client.GetEntriesParams{User: &user1},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, *byAlice.JSON200)
	byBob, e := c.GetEntriesWithResponse(
		x,
		&client.GetEntriesParams{User: &user2},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *byBob.JSON200)
	assert.String(t, "backup", (*byBob.JSON200)[0].Action)
	bySvc1, e := c.GetEntriesWithResponse(
		x,
		&client.GetEntriesParams{Service: &service1},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, *bySvc1.JSON200)
	bySvc2, e := c.GetEntriesWithResponse(
		x,
		&client.GetEntriesParams{Service: &service2},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *bySvc2.JSON200)
	byAliceSys2, e := c.GetEntriesWithResponse(
		x,
		&client.GetEntriesParams{User: &user1, System: &system2},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *byAliceSys2.JSON200)
	assert.String(t, "deploy", (*byAliceSys2.JSON200)[0].Action)
	byNobody, e := c.GetEntriesWithResponse(
		x,
		&client.GetEntriesParams{User: new("nobody")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, byNobody.StatusCode())
	assert.Count(t, 0, *byNobody.JSON200)
}
