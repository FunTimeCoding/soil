package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
	"time"
)

func TestLongStallClientWaitsLongForHeadersOnly(t *testing.T) {
	c := web.LongStallClient()
	transport, okay := c.Transport.(*http.Transport)
	assert.True(t, okay)
	assert.Any(t, time.Duration(0), c.Timeout)
	assert.Any(t, 30*time.Minute, transport.ResponseHeaderTimeout)
	assert.Any(t, 10*time.Second, transport.TLSHandshakeTimeout)
	assert.True(t, transport.ForceAttemptHTTP2)
	stall, _ := web.StallClient().Transport.(*http.Transport)
	assert.Any(t, 30*time.Second, stall.ResponseHeaderTimeout)
}
