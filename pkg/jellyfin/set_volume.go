package jellyfin

import "fmt"

func (c *Client) SetVolume(
	sessionIdentifier string,
	level int,
) error {
	return c.basic.Post(
		fmt.Sprintf("/Sessions/%s/Command", sessionIdentifier),
		nil,
		map[string]any{
			"Name":      "SetVolume",
			"Arguments": map[string]string{"Volume": fmt.Sprint(level)},
		},
	)
}
