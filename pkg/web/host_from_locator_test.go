package web

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestHostFromLocator(t *testing.T) {
	assert.String(
		t,
		"example.org",
		HostFromLocator(
			locator.New(
				constant.Example,
			).Port(constant.ListenPort).Path("/a").String(),
		),
	)
	assert.String(
		t,
		"example.org",
		HostFromLocator(locator.New(constant.Example).Path("/a").String()),
	)
}
