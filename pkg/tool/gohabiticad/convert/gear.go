package convert

import (
	"github.com/funtimecoding/soil/pkg/habitica/gear"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func Gear(g *gear.Gear) *server.GearResult {
	owned := make([]string, 0, len(g.Owned))

	for key, active := range g.Owned {
		if active {
			owned = append(owned, key)
		}
	}

	return &server.GearResult{Equipped: g.Equipped, Owned: owned}
}
