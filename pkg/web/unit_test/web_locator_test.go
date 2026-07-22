package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestParseLocator(t *testing.T) {
	actual := web.ParseLocator(
		locator.New(constant.Example).Port(constant.ListenPort).String(),
	)
	assert.String(t, "https://example.org:8080", actual.String())
	assert.String(t, "https", actual.Scheme)
	assert.String(t, "example.org:8080", actual.Host)
	assert.String(t, "8080", actual.Port())
	assert.String(t, "", actual.Path)
	assert.String(t, "", actual.RawQuery)
}

func TestHostFromLocator(t *testing.T) {
	assert.String(
		t,
		"example.org",
		web.HostFromLocator(
			locator.New(
				constant.Example,
			).Port(constant.ListenPort).Path("/a").String(),
		),
	)
	assert.String(
		t,
		"example.org",
		web.HostFromLocator(locator.New(constant.Example).Path("/a").String()),
	)
}

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

func TestHostPortFromLocatorSplit(t *testing.T) {
	host, port := web.HostPortFromLocatorSplit(
		locator.New(
			constant.Example,
		).Port(constant.ListenPort).Path("/a").String(),
	)
	assert.String(t, "example.org", host)
	assert.Integer(t, 8080, port)
}

func TestTrimScheme(t *testing.T) {
	assert.String(
		t,
		"localhost",
		web.TrimScheme(locator.New(constant.Localhost).String()),
	)
	assert.String(
		t,
		"localhost",
		web.TrimScheme(locator.New(constant.Localhost).Insecure().String()),
	)
}
