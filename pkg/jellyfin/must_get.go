package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/url"
)

func (c *Client) mustGet(
	path string,
	v url.Values,
	out any,
) {
	errors.PanicOnError(c.get(path, v, out))
}
