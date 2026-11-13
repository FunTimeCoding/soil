package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"testing"
)

func TestOutageIsAnsweredAsUnavailableWithoutCapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = refusedDial()
	m := memory.New()
	s := restClient(t, c, m)

	for range 2 {
		r, e := s.GetStatsWithResponse(context.Background())
		assert.FatalOnError(t, e)
		assert.NotNil(t, r.JSON503)
	}

	assert.Count(t, 0, m.Events())
}

func TestExpiredTokenIsAnsweredAsForbiddenWithoutCapture(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = forbidden.Format("habitica: %s", "token expired")
	m := memory.New()
	r, e := restClient(t, c, m).GetTasksWithResponse(context.Background(), nil)
	assert.FatalOnError(t, e)
	assert.NotNil(t, r.JSON403)
	assert.Count(t, 0, m.Events())
}

func TestConflictingCronIsAnsweredAsConflict(t *testing.T) {
	c := service_tester.NewSource(t)
	c.Failure = conflict.Format("habitica: %s", "cron already run")
	m := memory.New()
	r, e := restClient(t, c, m).RunCronWithResponse(context.Background())
	assert.FatalOnError(t, e)
	assert.NotNil(t, r.JSON409)
	assert.Count(t, 0, m.Events())
}
