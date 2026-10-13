package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStallClientPlain(t *testing.T) {
	p := stall_client.New().Build()
	defer p.Close()
	assert.True(t, p.Transport().TLSClientConfig == nil)
	assert.Duration(t, 30*time.Second, p.Transport().ResponseHeaderTimeout)
	assert.Duration(t, 0, p.Client().Timeout)
}

func TestStallClientInsecure(t *testing.T) {
	p := stall_client.New().Insecure().Build()
	defer p.Close()
	assert.True(t, p.Transport().TLSClientConfig.InsecureSkipVerify)
	assert.Duration(t, 30*time.Second, p.Transport().ResponseHeaderTimeout)
}

func TestStallClientLong(t *testing.T) {
	p := stall_client.New().Long().Build()
	defer p.Close()
	assert.Duration(t, 30*time.Minute, p.Transport().ResponseHeaderTimeout)
}

func TestStallClientTimeoutReachesBothViews(t *testing.T) {
	p := stall_client.New().Timeout(5 * time.Second).Build()
	defer p.Close()
	assert.Duration(t, 5*time.Second, p.Client().Timeout)
	assert.Duration(t, 5*time.Second, p.Halting().Timeout)
}

func TestStallClientViewsShareTheTransport(t *testing.T) {
	p := stall_client.New().Build()
	defer p.Close()
	assert.True(t, p.Client().Transport == p.Transport())
	assert.True(t, p.Halting().Transport == p.Transport())
}

func TestStallClientHaltingStopsAtTheRedirect(t *testing.T) {
	s := httptest.NewServer(
		http.RedirectHandler("/elsewhere", http.StatusFound),
	)
	defer s.Close()
	p := stall_client.New().Build()
	defer p.Close()
	r, e := p.Halting().Get(s.URL)
	assert.FatalOnError(t, e)
	defer errors.PanicClose(r.Body)
	assert.Integer(t, http.StatusFound, r.StatusCode)
}

func TestStallClientCloseDropsIdleConnections(t *testing.T) {
	s, opened := newCountingServer(t)
	p := stall_client.New().Build()
	defer p.Close()
	assert.String(t, "counted", web.GetString(p.Client(), s.URL))
	p.Close()
	assert.String(t, "counted", web.GetString(p.Client(), s.URL))
	assert.Integer(t, int64(2), opened.Load())
}
