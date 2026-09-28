package jellyfin

import (
	"fmt"
	"net/url"
)

func (c *Client) PlaybackCommand(
	sessionIdentifier string,
	command string,
	seekPositionTicks int64,
) error {
	v := url.Values{}

	if seekPositionTicks > 0 {
		v.Set("seekPositionTicks", fmt.Sprint(seekPositionTicks))
	}

	return c.post(
		fmt.Sprintf("/Sessions/%s/Playing/%s", sessionIdentifier, command),
		v,
		nil,
	)
}
