package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/worker"
	"testing"
)

func TestCommandOutageStaysQuietAndStopsAsking(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = unavailable.Format(
		"pgbackrest command unavailable: exit 125",
	)
	w, okay := s.Daemon().Workers()[1].(*worker.Worker)
	assert.True(t, okay)

	for range 3 {
		w.Poll()
	}

	assert.Count(t, 0, s.Reporter.Events())
	assert.Integer(t, 1, s.Source.Calls)
}

func TestPgbackrestFaultIsCapturedEachPoll(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = errors.New("pgbackrest info: exit status 31")
	w, okay := s.Daemon().Workers()[1].(*worker.Worker)
	assert.True(t, okay)
	w.Poll()
	w.Poll()
	assert.Count(t, 2, s.Reporter.Events())
}
