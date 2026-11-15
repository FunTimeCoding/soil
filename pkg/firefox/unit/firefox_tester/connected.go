package firefox_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/firefox"
	"github.com/funtimecoding/soil/pkg/firefox/unit/extension_peer"
	"testing"
)

func Connected(
	t *testing.T,
	c *firefox.Client,
	address string,
) *extension_peer.Extension {
	t.Helper()
	result := extension_peer.Connect(t, address)
	assert.Eventually(
		t,
		func() bool {
			_, e := c.Tabs()

			return !unavailable.Is(e)
		},
	)

	return result
}
