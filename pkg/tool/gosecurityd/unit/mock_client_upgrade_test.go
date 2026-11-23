package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"testing"
)

func TestMockListsSeededUpgradesLikeDaemon(t *testing.T) {
	m := mock_client.New(t)
	x := t.Context()
	m.SeedUpgrade(
		client.UpgradeGroup{
			Release:         "bookworm",
			Source:          "victor",
			Hosts:           []client.UpgradeHost{{Name: "foxtrot"}},
			Vulnerabilities: []string{"CVE-2099-30001"},
		},
	)
	m.SeedUpgrade(
		client.UpgradeGroup{
			Release: "bookworm",
			Source:  "linux",
			Kernel:  true,
			Hosts:   []client.UpgradeHost{{Name: "golf"}},
		},
	)
	all, e := m.ListUpgrades(x, nil)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, all.Groups)
	assert.Integer(t, 1, all.Hosts)
	kernel := true
	k, f := m.ListUpgrades(x, &client.ListUpgradesParams{Kernel: &kernel})
	assert.FatalOnError(t, f)
	assert.String(t, "linux", k.Groups[0].Source)
	trixie := "trixie"
	none, g := m.ListUpgrades(x, &client.ListUpgradesParams{Release: &trixie})
	assert.FatalOnError(t, g)
	assert.Count(t, 0, none.Groups)
	elsewhere := "golf"
	missed, j := m.ListUpgrades(
		x,
		&client.ListUpgradesParams{Search: &elsewhere},
	)
	assert.FatalOnError(t, j)
	assert.Count(t, 0, missed.Groups)
	one, h := m.GetUpgrade(x, "bookworm", "victor")
	assert.FatalOnError(t, h)
	assert.String(t, "foxtrot", one.Hosts[0].Name)
	_, i := m.GetUpgrade(x, "bookworm", "romeo")
	assert.True(t, not_found.Is(i))
}
