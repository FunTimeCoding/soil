package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestLocator(t *testing.T) {
	assert.String(
		t,
		"https://example.org",
		locator.New(constant.Example).String(),
	)
	assert.String(
		t,
		"http://example.org",
		locator.Stub().Insecure().Host(constant.Example).String(),
	)
	assert.String(
		t,
		"wss://example.org",
		locator.Stub().Insecure().Host(
			constant.Example,
		).Scheme(constant.SecureSocket).String(),
	)
}

func TestUser(t *testing.T) {
	assert.String(
		t,
		"https://u@example.org",
		locator.New(constant.Example).User("u").String(),
	)
	assert.String(
		t,
		"https://u:p@example.org",
		locator.New(constant.Example).UserPassword("u", "p").String(),
	)
}

func TestPort(t *testing.T) {
	assert.String(
		t,
		"https://example.org:8443",
		locator.New(constant.Example).Port(8443).String(),
	)
	assert.String(
		t,
		"https://example.org:8443/p",
		locator.New(constant.Example).Port(8443).Path("/p").String(),
	)
}

func TestPath(t *testing.T) {
	assert.String(
		t,
		"https://example.org/p",
		locator.New(constant.Example).Path("/p").String(),
	)
	assert.String(
		t,
		"https://example.org/b/p",
		locator.New(constant.Example).Base("/b").Path("/p").String(),
	)
	assert.String(
		t,
		"https://example.org/b/p",
		locator.New(constant.Example).Base("b").Path("/p").String(),
	)
	assert.String(
		t,
		"https://example.org/b/p",
		locator.New(constant.Example).Base("/b/").Path("/p").String(),
	)
	assert.String(
		t,
		"https://example.org/b/p",
		locator.New(constant.Example).Base("/b").Path("p").String(),
	)
	assert.String(
		t,
		"https://example.org/p/",
		locator.New(constant.Example).Path("/p/").Trail().String(),
	)
}

func TestParameter(t *testing.T) {
	assert.String(
		t,
		"https://example.org?a=1&b=2",
		locator.New(constant.Example).Add("a", "1").Add("b", "2").String(),
	)
	assert.String(
		t,
		"https://example.org/p?a=1&b=2",
		locator.New(constant.Example).Path(
			"/p",
		).Add("a", "1").Add("b", "2").String(),
	)
	assert.String(
		t,
		"https://example.org?a=1&a=2",
		locator.New(constant.Example).Add("a", "1").Add("a", "2").String(),
	)
	assert.String(
		t,
		"https://example.org?a=2",
		locator.New(constant.Example).Set("a", "1").Set("a", "2").String(),
	)
	assert.String(
		t,
		"https://example.org/p?a=b%3Dc",
		locator.New(constant.Example).Path("/p").Set("a", `b=c`).String(),
	)
}

func TestFragment(t *testing.T) {
	assert.String(
		t,
		"https://example.org#a",
		locator.New(constant.Example).Fragment("a").String(),
	)
	assert.String(
		t,
		"https://example.org/#/a?b=c%3Dd",
		locator.New(
			constant.Example,
		).Trail().Fragment("/a").FragmentSet("b", "c=d").String(),
	)
}
