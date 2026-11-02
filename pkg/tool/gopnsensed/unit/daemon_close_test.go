package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/mock_source"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/option"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/closing_source"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonClosesTheClientItBuilt(t *testing.T) {
	c := closing_source.New()
	o := option.New()
	o.Source = func() face.OpnsenseSource { return c }
	daemon.New(o).Close()
	assert.Integer(t, 1, c.Closed())
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	c := mock_source.New()
	before := goleak.IgnoreCurrent()
	daemon_tester.New(t, c).Close()
	goleak.VerifyNone(t, before)
}
