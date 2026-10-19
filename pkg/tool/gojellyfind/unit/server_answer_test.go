package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/server"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/rejecting_command"
	"testing"
)

func TestAnUnknownItemIsAnsweredNotFound(t *testing.T) {
	m := memory.New()
	r, e := server.New(mock_client.New(), m).GetItem(
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
	r, e := server.New(rejecting_command.New(), m).PlaybackCommand(
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
