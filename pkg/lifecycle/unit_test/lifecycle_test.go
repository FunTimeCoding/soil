package unit_test

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/mock_worker"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/location"
	"net/http"
	"testing"
)

func TestRunWorkerStartOrder(t *testing.T) {
	var log []string
	lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithWorker(mock_worker.New("a", &log)),
		lifecycle.WithWorker(mock_worker.New("b", &log)),
		lifecycle.WithWorker(mock_worker.New("c", &log)),
	).Run()
	assert.Count(t, 3, log)
	assert.String(t, "start:a", log[0])
	assert.String(t, "start:b", log[1])
	assert.String(t, "start:c", log[2])
}

func TestStopWorkerReverseOrder(t *testing.T) {
	var log []string
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithWorker(mock_worker.New("a", &log)),
		lifecycle.WithWorker(mock_worker.New("b", &log)),
		lifecycle.WithWorker(mock_worker.New("c", &log)),
	)
	l.Run()
	log = nil
	l.Stop()
	assert.Count(t, 3, log)
	assert.String(t, "stop:c", log[0])
	assert.String(t, "stop:b", log[1])
	assert.String(t, "stop:a", log[2])
}

func TestRunEmpty(t *testing.T) {
	lifecycle.New(logger.New(context.Background())).Run()
}

func TestStopEmpty(t *testing.T) {
	lifecycle.New(logger.New(context.Background())).Stop()
}

func TestRunServerResponds(t *testing.T) {
	p, n := system.ClaimPort()
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				"",
				func(m *http.ServeMux) {
					m.HandleFunc(
						location.Health,
						func(
							w http.ResponseWriter,
							_ *http.Request,
						) {
							w.WriteHeader(http.StatusOK)
						},
					)
				},
			).WithListener(n),
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, p)
	assert.HTTPStatus(
		t,
		fmt.Sprintf("http://localhost:%d/health", p),
		http.StatusOK,
	)
}

func TestStopServerShutsDown(t *testing.T) {
	p, n := system.ClaimPort()
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				"",
				func(m *http.ServeMux) {
					m.HandleFunc(
						location.Health,
						func(
							w http.ResponseWriter,
							_ *http.Request,
						) {
							w.WriteHeader(http.StatusOK)
						},
					)
				},
			).WithListener(n),
		),
	)
	l.Run()
	assert.Listen(t, p)
	l.Stop()
	assert.NotListen(t, p)
}

func TestRunServerMiddleware(t *testing.T) {
	p := system.FindUnusedPort(19300)
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				fmt.Sprintf(":%d", p),
				func(m *http.ServeMux) {
					m.HandleFunc(
						location.Health,
						func(
							w http.ResponseWriter,
							_ *http.Request,
						) {
							w.WriteHeader(http.StatusOK)
						},
					)
				},
			).WithMiddleware(
				func(h http.Handler) http.Handler {
					return http.HandlerFunc(
						func(
							w http.ResponseWriter,
							r *http.Request,
						) {
							w.Header().Set("X-Test", "middleware")
							h.ServeHTTP(w, r)
						},
					)
				},
			),
		),
	)
	l.Run()
	defer l.Stop()
	assert.Listen(t, p)
	r, e := http.Get(fmt.Sprintf("http://localhost:%d/health", p))
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(t, http.StatusOK, r.StatusCode)
	assert.String(t, "middleware", r.Header.Get("X-Test"))
}

func TestRunMixedOrder(t *testing.T) {
	var log []string
	p, n := system.ClaimPort()
	l := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithWorker(mock_worker.New("a", &log)),
		lifecycle.WithServer(
			server.New(
				"",
				func(m *http.ServeMux) {
					log = append(log, "start:server")
					m.HandleFunc(
						location.Health,
						func(
							w http.ResponseWriter,
							_ *http.Request,
						) {
							w.WriteHeader(http.StatusOK)
						},
					)
				},
			).WithListener(n),
		),
		lifecycle.WithWorker(mock_worker.New("b", &log)),
	)
	l.Run()
	assert.Count(t, 3, log)
	assert.String(t, "start:a", log[0])
	assert.String(t, "start:server", log[1])
	assert.String(t, "start:b", log[2])
	assert.Listen(t, p)
	log = nil
	l.Stop()
	assert.Count(t, 2, log)
	assert.String(t, "stop:b", log[0])
	assert.String(t, "stop:a", log[1])
}
