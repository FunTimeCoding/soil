package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/address"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestInternetType(t *testing.T) {
	assert.String(t, "IPv4", address.InternetType(constant.Loopback))
	assert.String(t, "IPv6", address.InternetType("::1"))
	assert.String(t, "IPv6", address.InternetType("2001:db8::1"))
	assert.String(t, "none", address.InternetType("invalid_ip"))
}
