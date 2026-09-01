package keepass

import "github.com/tobischo/gokeepasslib/v3"

func (c *Client) Walk(
	visit func(
		path string,
		group *gokeepasslib.Group,
		entry *gokeepasslib.Entry,
	),
) {
	for i := range c.database.Content.Root.Groups {
		walkGroup("", &c.database.Content.Root.Groups[i], visit)
	}
}
