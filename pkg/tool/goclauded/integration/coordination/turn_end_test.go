package coordination

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestTurnEndRecordsTheBoundary(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	assert.True(t, s.Store.GetSession(a.UUID).LastTurnEndAt == nil)
	response, e := s.RESTClient(t).PostTurnEndWithResponse(
		context.Background(),
		client.PostTurnEndJSONRequestBody{Session: a.UUID},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, response.StatusCode())
	assert.True(t, s.Store.GetSession(a.UUID).LastTurnEndAt != nil)
}

func TestTurnEndAfterPromptReadsIdle(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.CheckLive()
	s.Store.Advance(time.Minute)
	response, e := s.RESTClient(t).PostTurnEndWithResponse(
		context.Background(),
		client.PostTurnEndJSONRequestBody{Session: a.UUID},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, response.StatusCode())
	r := s.Store.GetSession(a.UUID)
	assert.True(t, r.LastTurnEndAt.After(*r.LastPromptAt))
}

func TestPromptAfterTurnEndReadsWorking(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	response, e := s.RESTClient(t).PostTurnEndWithResponse(
		context.Background(),
		client.PostTurnEndJSONRequestBody{Session: a.UUID},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, response.StatusCode())
	s.Store.Advance(time.Minute)
	a.CheckLive()
	r := s.Store.GetSession(a.UUID)
	assert.True(t, r.LastPromptAt.After(*r.LastTurnEndAt))
}

func TestTurnEndForUnknownSessionSucceeds(t *testing.T) {
	s := base.New(t)
	s.NewSession(t)
	response, e := s.RESTClient(t).PostTurnEndWithResponse(
		context.Background(),
		client.PostTurnEndJSONRequestBody{Session: uuid.New().String()},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 200, response.StatusCode())
}
