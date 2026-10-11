package technitium

import (
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/technitium/key"
	"github.com/funtimecoding/soil/pkg/technitium/response"
)

func (c *Client) Keys() ([]*key.Key, error) {
	var result response.Settings

	return result.Keys, c.get(constant.SettingsGetPath, &result)
}
