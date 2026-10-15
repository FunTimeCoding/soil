package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/funtimecoding/soil/pkg/web/unit/relay_tester"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStallClientDefaultPing(t *testing.T) {
	p := stall_client.New().Build()
	defer p.Close()
	assert.Duration(t, 30*time.Second, p.Transport().HTTP2.SendPingTimeout)
	assert.Duration(t, 15*time.Second, p.Transport().HTTP2.PingTimeout)
}

func TestStallClientDropsSilentConnection(t *testing.T) {
	s := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.WriteHeader(http.StatusNoContent)
			},
		),
	)
	s.EnableHTTP2 = true
	s.StartTLS()
	defer s.Close()
	r := relay_tester.New(t, s.Listener.Addr().String())
	p := stall_client.New().
		Insecure().
		Ping(100*time.Millisecond, 100*time.Millisecond).
		Timeout(2 * time.Second).
		Build()
	defer p.Close()
	target := join.Empty(constant.SecurePrefix, r.Address())
	first, e := p.Client().Get(target)
	assert.FatalOnError(t, e)
	errors.PanicClose(first.Body)
	assert.Integer(t, 2, first.ProtoMajor)
	r.Freeze()
	time.Sleep(time.Second)
	second, f := p.Client().Get(target)
	assert.FatalOnError(t, f)
	errors.PanicClose(second.Body)
	assert.Integer(t, http.StatusNoContent, second.StatusCode)
}
