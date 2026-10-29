package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/service_tester"
	"testing"
)

func TestDeleteOfUnknownEntryIsAnsweredNotFound(t *testing.T) {
	m := memory.New()
	s := server.New(service_tester.Service(t), m)
	r, e := s.DeleteEntry(
		context.Background(),
		generated.DeleteEntryRequestObject{Id: 9999},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.DeleteEntry404JSONResponse{Error: "entry not found: 9999"},
		r,
	)
	assert.Count(t, 0, m.Events())
}

func TestEntryWithoutActionIsAnsweredBadRequest(t *testing.T) {
	m := memory.New()
	s := server.New(service_tester.Service(t), m)
	r, e := s.PostEntry(
		context.Background(),
		generated.PostEntryRequestObject{
			Body: &generated.PostEntryJSONRequestBody{User: "alice"},
		},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.PostEntry400JSONResponse{Error: "action is required"},
		r,
	)
	assert.Count(t, 0, m.Events())
}

func TestStatusFaultIsCaptured(t *testing.T) {
	m := memory.New()
	v := service_tester.Service(t)
	s := server.New(v, m)
	v.Close()
	r, e := s.GetStatus(
		context.Background(),
		generated.GetStatusRequestObject{},
	)
	assert.FatalOnError(t, e)
	_, okay := r.(generated.GetStatus500JSONResponse)
	assert.True(t, okay)
	assert.Count(t, 1, m.Events())
}
