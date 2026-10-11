package technitium

import (
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"github.com/funtimecoding/soil/pkg/technitium/key"
	"net/url"
)

func (c *Client) SetKeys(v []*key.Key) error {
	return c.post(
		constant.SettingsSetPath,
		url.Values{constant.KeysParameter: {key.Encode(v)}},
	)
}
