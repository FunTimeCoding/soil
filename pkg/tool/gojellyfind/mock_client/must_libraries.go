package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/jellyfin/library"
)

func (c *Client) MustLibraries() []*library.Library {
	result, e := c.Libraries()
	errors.PanicOnError(e)

	return result
}
