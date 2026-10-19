package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generated "github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/server"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/rejecting_stat"
	"testing"
)

func TestAnInvalidStatIsAnsweredAsBadRequest(t *testing.T) {
	m := memory.New()
	s := server.New(rejecting_stat.New(), m)
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
