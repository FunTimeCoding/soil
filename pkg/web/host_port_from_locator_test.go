package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestHostPortFromLocator(t *testing.T) {
	assert.String(
		t,
		"example.org:8080",
		HostPortFromLocator(
			locator.New(
				constant.Example,
			).Port(constant.ListenPort).Path("/a").String(),
		),
	)
}
