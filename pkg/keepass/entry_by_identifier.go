package keepass

import (
	"fmt"
	"github.com/tobischo/gokeepasslib/v3"
)

func (c *Client) EntryByIdentifier(
	identifier string,
) (*gokeepasslib.Entry, *gokeepasslib.Group) {
	var foundEntry *gokeepasslib.Entry
	var foundGroup *gokeepasslib.Group
	c.Walk(
		func(
			_ string,
			group *gokeepasslib.Group,
			entry *gokeepasslib.Entry,
		) {
			if fmt.Sprintf("%x", entry.UUID) == identifier {
				foundEntry = entry
				foundGroup = group
			}
		},
	)

	return foundEntry, foundGroup
}
