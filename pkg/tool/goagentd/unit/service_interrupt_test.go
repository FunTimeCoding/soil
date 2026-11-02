package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/unit/stubborn_stream"
	"io"
	"os"
	"testing"
)

func TestAnInterruptBeforeTheLaunchKillsTheAgent(t *testing.T) {
	s := mock_stream.New(nil, true, nil)
	launching := make(chan struct{})
	release := make(chan struct{})
	o := service_tester.New(
		t,
		func(
			string,
			io.Writer,
			...string,
		) face.Stream {
			close(launching)
			<-release

			return s
		},
	)
	assert.FatalOnError(t, o.Service.Run("romeo"))
	<-launching
	o.Runner.Interrupt()
	close(release)
	service_tester.AwaitDone(t, o.Service)
	assert.Integer(t, 1, s.Killed())
}

func TestAnAgentAlreadyGoneIsNotReported(t *testing.T) {
	r := memory.New()
	o := service_tester.NewWith(
		t,
		service_tester.Fixed(stubborn_stream.New(os.ErrProcessDone)),
		r,
	)
	assert.FatalOnError(t, o.Service.Run("sierra"))
	o.Runner.Interrupt()
	service_tester.AwaitDone(t, o.Service)
	assert.Count(t, 0, r.Events())
}

func TestAKillThatFailsIsReported(t *testing.T) {
	r := memory.New()
	o := service_tester.NewWith(
		t,
		service_tester.Fixed(
			stubborn_stream.New(errors.New("operation not permitted")),
		),
		r,
	)
	assert.FatalOnError(t, o.Service.Run("tango"))
	o.Runner.Interrupt()
	service_tester.AwaitDone(t, o.Service)
	assert.Count(t, 1, r.Events())
}
