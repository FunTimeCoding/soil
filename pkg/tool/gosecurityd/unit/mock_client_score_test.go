package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"testing"
)

func TestMockNarrowsByScoreLikeDaemon(t *testing.T) {
	m := mock_client.New(t)
	x := t.Context()
	high := 9.8
	low := 5.0
	m.SeedSystemDetail(
		&client.SystemDetail{
			Summary: client.SystemSummary{Name: "tango", TopScore: high},
		},
	)
	m.SeedSystemDetail(
		&client.SystemDetail{
			Summary: client.SystemSummary{Name: "uniform", TopScore: low},
		},
	)
	m.SeedFindingGroup(
		&client.FindingGroup{Vulnerability: "CVE-2099-55001", Score: &high},
	)
	m.SeedFindingGroup(
		&client.FindingGroup{Vulnerability: "CVE-2099-55002", Score: &low},
	)
	m.SeedFindingGroup(&client.FindingGroup{Vulnerability: "CVE-2099-55003"})
	m.SeedUpgrade(
		client.UpgradeGroup{Release: "trixie", Source: "alpha", TopScore: high},
	)
	m.SeedUpgrade(
		client.UpgradeGroup{Release: "trixie", Source: "omega", TopScore: low},
	)
	critical := 9.0
	systems, e := m.ListSystems(x, &client.ListSystemsParams{Score: &critical})
	assert.FatalOnError(t, e)
	assert.Count(t, 1, systems)
	assert.String(t, "tango", systems[0].Name)
	groups, f := m.ListFindings(x, &client.ListFindingsParams{Score: &critical})
	assert.FatalOnError(t, f)
	assert.Count(t, 1, groups)
	assert.String(t, "CVE-2099-55001", groups[0].Vulnerability)
	all, g := m.ListFindings(x, nil)
	assert.FatalOnError(t, g)
	assert.Count(t, 3, all)
	upgrades, h := m.ListUpgrades(
		x,
		&client.ListUpgradesParams{Score: &critical},
	)
	assert.FatalOnError(t, h)
	assert.Count(t, 1, upgrades.Groups)
	assert.String(t, "alpha", upgrades.Groups[0].Source)
}
