package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MustSearchBlob(query string) []*gitlab.Blob {
	result, e := c.SearchBlob(query)
	errors.PanicOnError(e)

	return result
}
