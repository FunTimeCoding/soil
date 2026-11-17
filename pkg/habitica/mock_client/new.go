package mock_client

import (
	"github.com/funtimecoding/soil/pkg/habitica/gear"
	"github.com/funtimecoding/soil/pkg/habitica/statistic"
	"testing"
)

func New(t *testing.T) *Client {
	t.Helper()
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
	result := &Client{stats: stats, gear: g}
	t.Cleanup(result.Close)

	return result
}
