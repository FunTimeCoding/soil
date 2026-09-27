package bookmark

import "github.com/netbox-community/go-netbox/v4"

func New(b *netbox.Bookmark) *Bookmark {
	return &Bookmark{
		Identifier:       b.GetId(),
		Display:          b.GetDisplay(),
		ObjectType:       b.GetObjectType(),
		ObjectIdentifier: b.GetObjectId(),
		Link:             objectLink(b.GetObject()),
		Created:          b.GetCreated(),
		Raw:              b,
	}
}
