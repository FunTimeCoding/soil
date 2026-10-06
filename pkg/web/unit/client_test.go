package unit

import (
	"crypto/tls"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"testing"
	"time"
)

func TestClientIsBoundedLikeTheStallClient(t *testing.T) {
	transport, okay := web.Client().Transport.(*http.Transport)
	assert.True(t, okay)
	assert.Any(t, 30*time.Second, transport.ResponseHeaderTimeout)
	insecure, fine := web.InsecureClient().Transport.(*http.Transport)
	assert.True(t, fine)
	assert.Any(t, 30*time.Second, insecure.ResponseHeaderTimeout)
	assert.Any(
		t,
		&tls.Config{InsecureSkipVerify: true},
		insecure.TLSClientConfig,
	)
}
