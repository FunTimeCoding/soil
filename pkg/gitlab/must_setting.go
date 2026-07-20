package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) MustSetting() *gitlab.Settings {
	result, e := c.Setting()
	errors.PanicOnError(e)

	return result
}
