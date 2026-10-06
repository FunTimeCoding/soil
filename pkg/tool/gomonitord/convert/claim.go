package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store/claim"
)

func Claim(c *claim.Claim) server.Claim {
	return server.Claim{
		Item:      c.Item,
		Owner:     c.Owner,
		CreatedAt: c.CreatedAt,
	}
}
