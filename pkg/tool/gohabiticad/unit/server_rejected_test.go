package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	generated "github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/server"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"testing"
)

func TestAnInvalidStatIsAnsweredAsBadRequest(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = validation.New("%q is not a valid Stat.", "bravo")
	m := memory.New()
	s := server.New(service_tester.Service(t, c), m)
	r, e := s.AllocateStat(
		context.Background(),
		generated.AllocateStatRequestObject{Stat: "bravo"},
	)
	assert.FatalOnError(t, e)
	assert.Any(
		t,
		generated.AllocateStat400JSONResponse{
			Error: `"bravo" is not a valid Stat.`,
		},
		r,
	)
	assert.Count(t, 0, m.Events())
}
