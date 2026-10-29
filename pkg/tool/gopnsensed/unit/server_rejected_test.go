package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/server"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/rejecting_host"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/service_tester"
	"testing"
)

func TestARejectedHostIsAnsweredAsBadRequest(t *testing.T) {
	m := memory.New()
	s := server.New(service_tester.New(t, rejecting_host.New()), m)
	r, e := s.AddHost(
		context.Background(),
		generated.AddHostRequestObject{Body: &generated.HostRequest{}},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.AddHost400JSONResponse{
			Error: "host rejected: ip: invalid address",
		},
		r,
	)
	assert.Count(t, 0, m.Events())
}
