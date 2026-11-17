package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonSamplesWirelessWhenItAnswers(t *testing.T) {
	s := base.New(t)
	s.System.Answer(constant.Sudo, constant.WirelessSample)
	assert.Count(t, 4, s.Daemon().Workers())
}

func TestDaemonSkipsWirelessWhenItFails(t *testing.T) {
	s := base.New(t)
	s.System.Fail(constant.Sudo, errors.Format("exit", 1))
	assert.Count(t, 3, s.Daemon().Workers())
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := base.New(t)
	s.System.Serve(mock_stream.New(nil, true, errors.Format("killed", 9)))
	before := goleak.IgnoreCurrent()
	d := s.Daemon()
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	d.Close()
	goleak.VerifyNone(t, before)
}
