package habitica

import "github.com/funtimecoding/soil/pkg/strings/join"

func (c *Client) Equip(key string) error {
	return c.postDiscard(join.Empty("/user/equip/equipped/", key))
}
