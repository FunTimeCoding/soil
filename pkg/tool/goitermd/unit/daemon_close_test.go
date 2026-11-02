package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/face"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/option"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/closing_client"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonClosesTheClientItBuilt(t *testing.T) {
	c := closing_client.New()
	o := option.New()
	o.Source = func() face.ItermSource { return c }
	daemon.New(o).Close()
	assert.Integer(t, 1, c.Closed())
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	c := mock_client.New()
	before := goleak.IgnoreCurrent()
	daemon_tester.New(t, c).Close()
	goleak.VerifyNone(t, before)
}
