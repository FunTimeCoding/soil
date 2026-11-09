package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	c := mock_client.New()
	before := goleak.IgnoreCurrent()
	daemon_tester.New(t, c).Close()
	goleak.VerifyNone(t, before)
}
