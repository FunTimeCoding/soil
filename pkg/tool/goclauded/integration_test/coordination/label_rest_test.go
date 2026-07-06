//go:build local

package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func labelBody(
	key string,
	value string,
	from string,
) client.LabelRequest {
	return client.LabelRequest{Key: key, Value: &value, From: &from}
}

func TestRestLabelSetUpdateRemove(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
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
