package netbox

import "github.com/funtimecoding/soil/pkg/netbox/cache"

func (c *Client) Forget() {
	c.cache = cache.New()
}
