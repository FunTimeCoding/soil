package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustAdminMode(on bool) *gitlab.Settings {
	result, e := c.AdminMode(on)
	errors.PanicOnError(e)

	return result
}
