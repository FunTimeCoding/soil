package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/base"
	"testing"
)

func TestTheClientRoundTripsAnEntry(t *testing.T) {
	c := base.New(t).MaintenanceClient(t)
	x := t.Context()
	system := "worker1"
	created, e := c.PostEntry(
		x,
		client.PostEntryRequest{
			Action: "restart",
			User:   "alice",
			System: &system,
		},
	)
	assert.FatalOnError(t, e)
	assert.String(t, "restart", created.Action)
	listed, f := c.GetEntries(x, &client.GetEntriesParams{System: &system})
	assert.FatalOnError(t, f)
	assert.Count(t, 1, listed)
	updated, g := c.UpdateEntry(
		x,
		created.Id,
		client.PostEntryRequest{Action: "reboot", User: "alice"},
	)
	assert.FatalOnError(t, g)
	assert.String(t, "reboot", updated.Action)
	assert.FatalOnError(t, c.DeleteEntry(x, created.Id))
	status, h := c.GetStatus(x)
	assert.FatalOnError(t, h)
	assert.Integer(t, 0, status.TotalEntries)
}

func TestTheClientClassifiesRefusals(t *testing.T) {
	c := base.New(t).MaintenanceClient(t)
	x := t.Context()
	assert.True(t, not_found.Is(c.DeleteEntry(x, 9999)))
	_, e := c.UpdateEntry(
		x,
		9999,
		client.PostEntryRequest{Action: "x", User: "y"},
	)
	assert.True(t, not_found.Is(e))
	_, f := c.PostEntry(x, client.PostEntryRequest{User: "alice"})
	assert.True(t, validation.Is(f))
}
