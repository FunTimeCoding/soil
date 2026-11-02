package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/rejecting_command"
	"testing"
)

func TestAnUnknownItemIsAnsweredNotFound(t *testing.T) {
	m := memory.New()
	r, e := restClient(t, mock_client.New(), m).GetItemWithResponse(
		context.Background(),
		"charlie",
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON404)
	assert.String(t, "item not found: charlie", r.JSON404.Error)
	assert.Count(t, 0, m.Events())
}

func TestARejectedCommandIsAnsweredAsBadRequest(t *testing.T) {
	m := memory.New()
	r, e := restClient(
		t,
		rejecting_command.New(),
		m,
	).PlaybackCommandWithResponse(
		context.Background(),
		"delta",
		client.PlaybackCommandJSONRequestBody{Command: "Alfa"},
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON400)
	assert.String(t, "One or more validation errors occurred.", r.JSON400.Error)
	assert.Count(t, 0, m.Events())
}

func TestPlayWithoutItemsIsAnsweredAsBadRequest(t *testing.T) {
	m := memory.New()
	r, e := restClient(t, mock_client.New(), m).PlayWithResponse(
		context.Background(),
		"delta",
		client.PlayJSONRequestBody{ItemIds: []string{}},
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON400)
	assert.String(t, "item_ids is required", r.JSON400.Error)
	assert.Count(t, 0, m.Events())
}

func TestAVolumeOutOfRangeIsAnsweredAsBadRequest(t *testing.T) {
	m := memory.New()
	r, e := restClient(t, mock_client.New(), m).SetVolumeWithResponse(
		context.Background(),
		"delta",
		client.SetVolumeJSONRequestBody{Level: 101},
	)
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r.JSON400)
	assert.String(t, "level must be between 0 and 100", r.JSON400.Error)
	assert.Count(t, 0, m.Events())
}
