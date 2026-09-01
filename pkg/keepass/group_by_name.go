package keepass

import "github.com/tobischo/gokeepasslib/v3"

func (c *Client) GroupByName(name string) *gokeepasslib.Group {
	return GroupByNameRecursive(c.database.Content.Root.Groups, name)
}
