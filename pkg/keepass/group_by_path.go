package keepass

import (
	"github.com/tobischo/gokeepasslib/v3"
	"strings"
)

func (c *Client) GroupByPath(path string) *gokeepasslib.Group {
	segments := strings.Split(path, "/")
	groups := c.database.Content.Root.Groups
	var found *gokeepasslib.Group

	for _, segment := range segments {
		found = nil

		for i := range groups {
			if groups[i].Name == segment {
				found = &groups[i]

				break
			}
		}

		if found == nil {
			return nil
		}

		groups = found.Groups
	}

	return found
}
