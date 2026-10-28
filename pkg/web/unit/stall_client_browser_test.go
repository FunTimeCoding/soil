package unit

import (
	"crypto/tls"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"
)

func TestStallClientBrowserSendsBrowserHello(t *testing.T) {
	var hello *tls.ClientHelloInfo
	s := httptest.NewUnstartedServer(
		http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}),
	)
	s.EnableHTTP2 = true
	s.TLS = &tls.Config{
		GetConfigForClient: func(
			i *tls.ClientHelloInfo,
		) (*tls.Config, error) {
			hello = i

			return nil, nil
		},
	}
	s.StartTLS()
	defer s.Close()
	p := stall_client.New().Browser().Insecure().Build()
	defer p.Close()
	r, e := p.Client().Get(s.URL)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.String(t, "HTTP/1.1", r.Proto)
	assert.Strings(t, []string{"http/1.1"}, hello.SupportedProtos)
	assert.True(
		t,
		slices.ContainsFunc(
			hello.CipherSuites,
			func(c uint16) bool { return c&0x0f0f == 0x0a0a },
		),
	)
}
