package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestHostPortFromLocator(t *testing.T) {
	assert.String(
		t,
		"example.org:8080",
		web.HostPortFromLocator(
			locator.New(
				constant.Example,
			).Port(constant.ListenPort).Path("/a").String(),
		),
	)
}
