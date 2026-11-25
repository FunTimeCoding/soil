package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
	"time"
)

func TestRestRosterWindow(t *testing.T) {
	s := base.New(t)
	a := s.NewSession()
	a.Announce(a.Name(), "earlier")
	s.Store.Advance(2 * time.Hour)
	b := s.NewSession()
	b.Announce(b.Name(), "later")
	c := s.RESTClient()
	current, e := c.GetRosterWithResponse(
		t.Context(),
		&client.GetRosterParams{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, current.StatusCode())
	assert.Count(t, 1, current.JSON200.Entries)
	assert.String(t, b.Name(), current.JSON200.Entries[0].Callsign)
	assert.String(t, b.UUID, current.JSON200.Entries[0].Identifier)
	assert.True(t, current.JSON200.Entries[0].OnRoster)
	day := "24h"
	wide, f := c.GetRosterWithResponse(
		t.Context(),
		&client.GetRosterParams{SeenWithin: &day},
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 2, wide.JSON200.Entries)
	state := map[string]bool{}

	for _, v := range wide.JSON200.Entries {
		state[v.Callsign] = v.OnRoster
	}

	assert.False(t, state[a.Name()])
	assert.True(t, state[b.Name()])
	s.Store.Advance(25 * time.Hour)
	gone, g := c.GetRosterWithResponse(
		t.Context(),
		&client.GetRosterParams{SeenWithin: &day},
	)
	assert.FatalOnError(t, g)
	assert.Count(t, 0, gone.JSON200.Entries)
}

func TestRestRosterUnreadableWindow(t *testing.T) {
	soon := "soon"
	r, e := base.New(t).RESTClient().GetRosterWithResponse(
		t.Context(),
		&client.GetRosterParams{SeenWithin: &soon},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 400, r.StatusCode())
	assert.StringContains(
		t,
		"seen_within is not a duration (24h, 7d): soon",
		r.JSON400.Error,
	)
}
