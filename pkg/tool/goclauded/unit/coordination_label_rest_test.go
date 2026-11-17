package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/label_tester"
	"testing"
)

func TestRestLabelSetUpdateRemove(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	set, e := a.GeneratedClient.PostSessionLabelWithResponse(
		a.Context,
		a.UUID,
		label_tester.Body("environment", "staging", "reconciler"),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, set.StatusCode())
	assert.StringContains(t, "(unset)→staging", set.JSON200.Change)
	update, e := a.GeneratedClient.PostSessionLabelWithResponse(
		a.Context,
		a.UUID,
		label_tester.Body("environment", "production", "reconciler"),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, update.StatusCode())
	assert.StringContains(t, "staging→production", update.JSON200.Change)
	remove, e := a.GeneratedClient.PostSessionLabelWithResponse(
		a.Context,
		a.UUID,
		label_tester.Body("environment", "", "reconciler"),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, remove.StatusCode())
	assert.StringContains(t, "production→ (unset)", remove.JSON200.Change)
}

func TestRestLabelReadsBack(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "working")
	a.CheckLive()
	_, e := a.GeneratedClient.PostSessionLabelWithResponse(
		a.Context,
		a.UUID,
		label_tester.Body("environment", "staging", "reconciler"),
	)
	assert.FatalOnError(t, e)
	read, e := a.GeneratedClient.GetSessionLabelWithResponse(a.Context, a.UUID)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, read.StatusCode())
	assert.Any(
		t,
		[]client.LabelEntry{{Key: "environment", Value: "staging"}},
		read.JSON200.Labels,
	)
}

func TestRestLabelReadIsEmptyForUnknownSession(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	read, e := a.GeneratedClient.GetSessionLabelWithResponse(
		a.Context,
		"nonexistent",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, read.StatusCode())
	assert.Count(t, 0, read.JSON200.Labels)
}
