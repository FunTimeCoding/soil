package firefox_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/firefox"
	"github.com/funtimecoding/soil/pkg/firefox/unit/extension_peer"
)

func (o *Tester) Connected(
	c *firefox.Client,
	address string,
) *extension_peer.Extension {
	o.t.Helper()
	result := extension_peer.Connect(o.t, address)
	assert.Eventually(
		o.t,
		func() bool {
			_, e := c.Tabs()

			return !unavailable.Is(e)
		},
	)

	return result
}
