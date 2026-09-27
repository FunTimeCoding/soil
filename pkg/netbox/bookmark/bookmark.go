package bookmark

import (
	"github.com/netbox-community/go-netbox/v4"
	"time"
)

type Bookmark struct {
	Identifier       int32
	Display          string
	ObjectType       string
	ObjectIdentifier int64
	Link             string
	Created          time.Time
	Raw              *netbox.Bookmark
}
