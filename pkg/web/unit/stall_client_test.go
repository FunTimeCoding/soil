package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"testing"
	"time"
)

func TestStallClientBoundsWaitingNotBodies(t *testing.T) {
	p := stall_client.New().Build()
	defer p.Close()
	assert.Any(t, time.Duration(0), p.Client().Timeout)
	assert.Any(t, 30*time.Second, p.Transport().ResponseHeaderTimeout)
	assert.Any(t, 10*time.Second, p.Transport().TLSHandshakeTimeout)
	assert.True(t, p.Transport().ForceAttemptHTTP2)
}
