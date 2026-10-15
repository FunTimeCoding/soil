package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/funtimecoding/soil/pkg/web/unit/observation_tester"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStallClientObservesResponseAndRefusal(t *testing.T) {
	o := observation_tester.New()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.WriteHeader(http.StatusTeapot)
			},
		),
	)
	p := stall_client.New().Observe(o.Histogram).Build()
	defer p.Close()
	v, e := p.Client().Get(s.URL)
	assert.FatalOnError(t, e)
	errors.PanicClose(v.Body)
	s.Close()
	_, e = p.Client().Get(s.URL)
	assert.True(t, e != nil)
	assert.Any(
		t,
		map[string]uint64{"418 GET": 1, "unreachable GET": 1},
		o.Codes(),
	)
	assert.Strings(t, []string{s.Listener.Addr().String()}, o.Hosts())
}

func TestStallClientObservesDefaultPort(t *testing.T) {
	o := observation_tester.New()
	p := stall_client.New().Observe(o.Histogram).Build()
	defer p.Close()
	x, cancel := context.WithCancel(context.Background())
	cancel()
	_, e := p.Client().Do(getWithContext(x, "https://charlie.example/"))
	assert.True(t, e != nil)
	_, e = p.Client().Do(getWithContext(x, "http://charlie.example/"))
	assert.True(t, e != nil)
	assert.Strings(
		t,
		[]string{"charlie.example:443", "charlie.example:80"},
		o.Hosts(),
	)
}

func TestStallClientObservesTimeoutAndCancel(t *testing.T) {
	o := observation_tester.New()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				_ http.ResponseWriter,
				q *http.Request,
			) {
				<-q.Context().Done()
			},
		),
	)
	defer s.Close()
	p := stall_client.New().Observe(o.Histogram).Build()
	defer p.Close()
	x, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	_, e := p.Client().Do(getWithContext(x, s.URL))
	assert.True(t, e != nil)
	y, stop := context.WithCancel(context.Background())
	time.AfterFunc(100*time.Millisecond, stop)
	_, e = p.Client().Do(getWithContext(y, s.URL))
	assert.True(t, e != nil)
	assert.Any(
		t,
		map[string]uint64{"timeout GET": 1, "canceled GET": 1},
		o.Codes(),
	)
}

func TestStallClientObservesDroppedConnection(t *testing.T) {
	o := observation_tester.New()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				c, _, e := http.NewResponseController(w).Hijack()
				errors.PanicOnError(e)
				errors.PanicOnError(c.(*net.TCPConn).SetLinger(0))
				errors.PanicClose(c)
			},
		),
	)
	defer s.Close()
	p := stall_client.New().Observe(o.Histogram).Build()
	defer p.Close()
	_, e := p.Client().Get(s.URL)
	assert.True(t, e != nil)
	assert.Any(t, map[string]uint64{"dropped GET": 1}, o.Codes())
}
