package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/rejecting_command"
	"testing"
)

func TestUnknownItemIsAnsweredNotFound(t *testing.T) {
	s := base.New(t)
	r, e := s.GeneratedClient().GetItemWithResponse(
		context.Background(),
		"charlie",
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON404)
	assert.String(t, "item not found: charlie", r.JSON404.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestRejectedCommandIsAnsweredAsBadRequest(t *testing.T) {
	s := base.New(t, base.WithSource(rejecting_command.New(t)))
	r, e := s.GeneratedClient().PlaybackCommandWithResponse(
		context.Background(),
		"delta",
		client.PlaybackCommandJSONRequestBody{Command: "Alfa"},
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON400)
	assert.String(t, "One or more validation errors occurred.", r.JSON400.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestPlayWithoutItemsIsAnsweredAsBadRequest(t *testing.T) {
	s := base.New(t)
	r, e := s.GeneratedClient().PlayWithResponse(
		context.Background(),
		"delta",
		client.PlayJSONRequestBody{ItemIds: []string{}},
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON400)
	assert.String(t, "item_ids is required", r.JSON400.Error)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestVolumeOutOfRangeIsAnsweredAsBadRequest(t *testing.T) {
	s := base.New(t)
	r, e := s.GeneratedClient().SetVolumeWithResponse(
		context.Background(),
		"delta",
		client.SetVolumeJSONRequestBody{Level: 101},
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON400)
	assert.String(t, "level must be between 0 and 100", r.JSON400.Error)
	assert.Count(t, 0, s.Reporter.Events())
}
