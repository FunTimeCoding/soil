package mock_client

import (
	"github.com/funtimecoding/soil/pkg/habitica/gear"
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
)

func New() *Client {
	stats := statistic.Stub()
	stats.HP = 50
	stats.MP = 30
	stats.XP = 100
	stats.GP = 10
	stats.Level = 5
	stats.Class = "warrior"
	g := gear.Stub()
	g.Equipped = map[string]string{}
	g.Owned = map[string]bool{}

	return &Client{stats: stats, gear: g}
}
