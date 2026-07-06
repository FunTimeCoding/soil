package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestHostPortFromLocatorSplit(t *testing.T) {
	host, port := HostPortFromLocatorSplit(
		locator.New(
			constant.Example,
		).Port(constant.ListenPort).Path("/a").String(),
	)
	assert.String(t, "example.org", host)
	assert.Integer(t, 8080, port)
}
