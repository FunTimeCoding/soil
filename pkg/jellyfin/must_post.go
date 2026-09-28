package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/url"
)

func (c *Client) mustPost(
	path string,
	v url.Values,
	body any,
) {
	errors.PanicOnError(c.post(path, v, body))
}
