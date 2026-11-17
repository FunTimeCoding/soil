package unit

import (
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	stream := mock_stream.New(nil, false, nil)
	t.Cleanup(stream.Close)
	s := service_tester.Fixed(stream)
	before := goleak.IgnoreCurrent()
	d := base.New(t, s).Daemon()
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
