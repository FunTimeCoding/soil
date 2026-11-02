package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	c := mock_client.New()
	before := goleak.IgnoreCurrent()
	daemon_tester.New(t, c).Close()
	goleak.VerifyNone(t, before)
}
