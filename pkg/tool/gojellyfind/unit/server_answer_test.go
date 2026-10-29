package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/server"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/rejecting_command"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/service_tester"
	"testing"
)

func TestAnUnknownItemIsAnsweredNotFound(t *testing.T) {
	m := memory.New()
	r, e := server.New(service_tester.New(t, mock_client.New()), m).GetItem(
		context.Background(),
		generated.GetItemRequestObject{Identifier: "charlie"},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.GetItem404JSONResponse{Error: "item not found: charlie"},
		r,
	)
	assert.Count(t, 0, m.Events())
}

func TestARejectedCommandIsAnsweredAsBadRequest(t *testing.T) {
	m := memory.New()
	r, e := server.New(
		service_tester.New(t, rejecting_command.New()),
		m,
	).PlaybackCommand(
		context.Background(),
		generated.PlaybackCommandRequestObject{
			Identifier: "delta",
			Body: &generated.PlaybackCommandJSONRequestBody{
				Command: "Alfa",
			},
		},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.PlaybackCommand400JSONResponse{
			Error: "One or more validation errors occurred.",
		},
		r,
	)
	assert.Count(t, 0, m.Events())
}

func TestPlayWithoutItemsIsAnsweredAsBadRequest(t *testing.T) {
	m := memory.New()
	r, e := server.New(service_tester.New(t, mock_client.New()), m).Play(
		context.Background(),
		generated.PlayRequestObject{
			Identifier: "delta",
			Body:       &generated.PlayJSONRequestBody{ItemIds: []string{}},
		},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.Play400JSONResponse{Error: "item_ids is required"},
		r,
	)
	assert.Count(t, 0, m.Events())
}

func TestAVolumeOutOfRangeIsAnsweredAsBadRequest(t *testing.T) {
	m := memory.New()
	r, e := server.New(service_tester.New(t, mock_client.New()), m).SetVolume(
		context.Background(),
		generated.SetVolumeRequestObject{
			Identifier: "delta",
			Body:       &generated.SetVolumeJSONRequestBody{Level: 101},
		},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.SetVolume400JSONResponse{
			Error: "level must be between 0 and 100",
		},
		r,
	)
	assert.Count(t, 0, m.Events())
}
