package jellyfin

import (
	"fmt"
	"net/url"
	"strings"
)

func (c *Client) Play(
	sessionIdentifier string,
	itemIDs []string,
	playCommand string,
	startPositionTicks int64,
) error {
	v := url.Values{}
	v.Set("playCommand", playCommand)
	v.Set("itemIds", strings.Join(itemIDs, ","))

	if startPositionTicks > 0 {
		v.Set("startPositionTicks", fmt.Sprint(startPositionTicks))
	}

	return c.basic.Post(
		fmt.Sprintf("/Sessions/%s/Playing", sessionIdentifier),
		v,
		nil,
	)
}
