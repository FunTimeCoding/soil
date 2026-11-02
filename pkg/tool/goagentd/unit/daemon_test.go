package unit

import (
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	s := service_tester.Fixed(mock_stream.New(nil, false, nil))
	before := goleak.IgnoreCurrent()
	d := daemon_tester.New(t, s)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}
