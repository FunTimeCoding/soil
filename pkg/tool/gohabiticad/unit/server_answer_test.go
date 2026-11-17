package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOutageIsAnsweredAsUnavailableWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	c := s.GeneratedClient()

	for range 2 {
		r, e := c.GetStatsWithResponse(context.Background())
		assert.FatalOnError(t, e)
		assert.NotNil(t, r.JSON503)
	}

	assert.Count(t, 0, s.Reporter.Events())
}

func TestExpiredTokenIsAnsweredAsForbiddenWithoutCapture(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = forbidden.Format("habitica: %s", "token expired")
	r, e := s.GeneratedClient().GetTasksWithResponse(context.Background(), nil)
	assert.FatalOnError(t, e)
	assert.NotNil(t, r.JSON403)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestConflictingCronIsAnsweredAsConflict(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = conflict.Format("habitica: %s", "cron already run")
	r, e := s.GeneratedClient().RunCronWithResponse(context.Background())
	assert.FatalOnError(t, e)
	assert.NotNil(t, r.JSON409)
	assert.Count(t, 0, s.Reporter.Events())
}
