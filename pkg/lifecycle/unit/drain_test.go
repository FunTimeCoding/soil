package unit

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/constant"
	"github.com/funtimecoding/soil/pkg/lifecycle/mock_drainer"
	"github.com/funtimecoding/soil/pkg/lifecycle/mock_worker"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/lifecycle/worker"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"net/http"
	"testing"
	"time"
)

func TestStopDrainerFirstAndOnce(t *testing.T) {
	var log []string
	l := lifecycle.New(logger.New(context.Background())).
		WithWorker(mock_worker.New("a", &log)).
		WithWorker(mock_drainer.New("d", &log)).
		WithWorker(mock_worker.New("b", &log))
	l.Run()
	log = nil
	l.Close()
	assert.Strings(t, []string{"stop:d", "stop:b", "stop:a"}, log)
}

func TestStopInterruptsHeldDrainerAtGrace(t *testing.T) {
	var log []string
	l := lifecycle.New(logger.New(context.Background())).
		WithGrace(constant.InterruptMargin + 10*time.Millisecond).
		WithWorker(mock_drainer.New("d", &log).Hold())
	l.Run()
	log = nil
	start := time.Now()
	l.Close()
	assert.Less(t, time.Second, time.Since(start))
	assert.Strings(t, []string{"stop:d", "interrupt:d"}, log)
}

func TestDrainReturnsWhenFinished(t *testing.T) {
	var log []string
	w := worker.New(mock_drainer.New("d", &log))
	w.Close()
	w.Drain(time.Now().Add(time.Second))
	assert.Strings(t, []string{"stop:d"}, log)
}

func TestDrainInterruptsAtDeadline(t *testing.T) {
	var log []string
	w := worker.New(mock_drainer.New("d", &log).Hold())
	w.Close()
	w.Drain(time.Now().Add(10 * time.Millisecond))
	assert.Strings(t, []string{"stop:d", "interrupt:d"}, log)
}

func TestDrainIgnoresPlainWorker(t *testing.T) {
	var log []string
	w := worker.New(mock_worker.New("a", &log))
	w.Close()
	w.Drain(time.Now().Add(time.Hour))
	assert.Strings(t, []string{"stop:a"}, log)
}

func TestCloseReturnsAfterIgnoredInterrupt(t *testing.T) {
	var log []string
	l := lifecycle.New(logger.New(context.Background())).
		WithGrace(constant.InterruptMargin).
		WithWorker(mock_drainer.New("d", &log).Hold().Ignore())
	l.Run()
	log = nil
	start := time.Now()
	l.Close()
	assert.Greater(t, 4*time.Second, time.Since(start))
	assert.Less(t, 6*time.Second, time.Since(start))
	assert.Strings(t, []string{"stop:d", "interrupt:d"}, log)
}

func TestDrainersShareOneDeadline(t *testing.T) {
	var log []string
	l := lifecycle.New(logger.New(context.Background())).
		WithGrace(constant.InterruptMargin + 500*time.Millisecond).
		WithWorker(mock_drainer.New("a", &log).Hold()).
		WithWorker(mock_drainer.New("b", &log).Hold())
	l.Run()
	log = nil
	start := time.Now()
	l.Close()
	assert.Less(t, 800*time.Millisecond, time.Since(start))
	assert.Strings(
		t,
		[]string{"stop:a", "stop:b", "interrupt:a", "interrupt:b"},
		log,
	)
}

func TestServerShutdownCountsAgainstGrace(t *testing.T) {
	var log []string
	arrived := make(chan struct{})
	release := make(chan struct{})
	defer close(release)
	p, n := system.ClaimPort()
	l := lifecycle.New(logger.New(context.Background())).
		WithGrace(constant.InterruptMargin + 3*time.Second).
		WithWorker(mock_drainer.New("d", &log).Hold()).
		WithServer(
			server.New(
				identity.Example(),
				"",
				func(m *http.ServeMux) {
					m.HandleFunc(
						"/hold",
						func(
							_ http.ResponseWriter,
							_ *http.Request,
						) {
							close(arrived)
							<-release
						},
					)
				},
			).WithListener(n),
		)
	l.Run()
	o := stall_client.New().Build()
	defer o.Close()
	assert.Listen(t, p)
	go func() {
		response, e := o.Client().Get(
			fmt.Sprintf("http://localhost:%d/hold", p),
		)

		if e != nil {
			return
		}

		errors.PanicClose(response.Body)
	}()
	<-arrived
	log = nil
	start := time.Now()
	l.Close()
	assert.Greater(t, 4*time.Second, time.Since(start))
	assert.Less(t, 7*time.Second, time.Since(start))
	assert.Strings(t, []string{"stop:d", "interrupt:d"}, log)
}
