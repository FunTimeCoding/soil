package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
	gomaintlogd "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/service_tester"
	"testing"
)

func TestDeleteOfUnknownEntryIsAnsweredNotFound(t *testing.T) {
	m := memory.New()
	s := gomaintlogd.New(service_tester.Service(t), m)
	r, e := s.DeleteEntry(
		context.Background(),
		server.DeleteEntryRequestObject{Id: 9999},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		server.DeleteEntry404JSONResponse{Error: "entry not found: 9999"},
		r,
	)
	assert.Count(t, 0, m.Events())
}

func TestEntryWithoutActionIsAnsweredBadRequest(t *testing.T) {
	m := memory.New()
	s := gomaintlogd.New(service_tester.Service(t), m)
	r, e := s.PostEntry(
		context.Background(),
		server.PostEntryRequestObject{
			Body: &server.PostEntryJSONRequestBody{User: "alice"},
		},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		server.PostEntry400JSONResponse{Error: "action is required"},
		r,
	)
	assert.Count(t, 0, m.Events())
}

func TestStatusFaultIsCaptured(t *testing.T) {
	m := memory.New()
	o := service_tester.New(t)
	s := gomaintlogd.New(o.Service, m)
	o.Close()
	r, e := s.GetStatus(context.Background(), server.GetStatusRequestObject{})
	assert.FatalOnError(t, e)
	_, okay := r.(server.GetStatus500JSONResponse)
	assert.True(t, okay)
	assert.Count(t, 1, m.Events())
}
