package address

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestInternetType(t *testing.T) {
	assert.String(t, "IPv4", InternetType("127.0.0.1"))
	assert.String(t, "IPv6", InternetType("::1"))
	assert.String(t, "IPv6", InternetType("2001:db8::1"))
	assert.String(t, "none", InternetType("invalid_ip"))
}
