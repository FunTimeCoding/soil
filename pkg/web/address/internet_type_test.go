package address

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestInternetType(t *testing.T) {
	assert.String(t, "IPv4", InternetType(constant.Loopback))
	assert.String(t, "IPv6", InternetType("::1"))
	assert.String(t, "IPv6", InternetType("2001:db8::1"))
	assert.String(t, "none", InternetType("invalid_ip"))
}
