package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
)

func TestRestLabelSetUpdateRemove(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.CheckLive()
	set, e := a.RestClient.PostSessionLabelWithResponse(
		a.Context,
		a.UUID,
		labelBody("environment", "staging", "reconciler"),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, set.StatusCode())
	assert.StringContains(t, "(unset)→staging", set.JSON200.Change)
	update, e := a.RestClient.PostSessionLabelWithResponse(
		a.Context,
		a.UUID,
		labelBody("environment", "production", "reconciler"),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, update.StatusCode())
	assert.StringContains(t, "staging→production", update.JSON200.Change)
	remove, e := a.RestClient.PostSessionLabelWithResponse(
		a.Context,
		a.UUID,
		labelBody("environment", "", "reconciler"),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, remove.StatusCode())
	assert.StringContains(t, "production→ (unset)", remove.JSON200.Change)
}

func TestRestLabelReadsBack(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "working")
	a.CheckLive()
	_, e := a.RestClient.PostSessionLabelWithResponse(
		a.Context,
		a.UUID,
		labelBody("environment", "staging", "reconciler"),
	)
	assert.FatalOnError(t, e)
	read, e := a.RestClient.GetSessionLabelWithResponse(a.Context, a.UUID)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, read.StatusCode())
	assert.Any(
		t,
		[]client.LabelEntry{{Key: "environment", Value: "staging"}},
		read.JSON200.Labels,
	)
}

func TestRestLabelReadIsEmptyForAnUnknownSession(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	read, e := a.RestClient.GetSessionLabelWithResponse(
		a.Context,
		"nonexistent",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, read.StatusCode())
	assert.Count(t, 0, read.JSON200.Labels)
}
