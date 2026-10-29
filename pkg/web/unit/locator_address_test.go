package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestAddressNamesTheHostAndPort(t *testing.T) {
	assert.String(
		t,
		"example.org:8443",
		locator.New(constant.Example).Port(8443).Path("/p").Address(),
	)
	assert.String(t, "example.org", locator.New(constant.Example).Address())
}

func TestAddressLeavesTheCredentialsOut(t *testing.T) {
	assert.String(
		t,
		"example.org",
		locator.New(constant.Example).UserPassword("u", "p").Address(),
	)
}
