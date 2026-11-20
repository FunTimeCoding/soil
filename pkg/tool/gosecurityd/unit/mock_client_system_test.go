package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"testing"
)

func TestMockStoresSnapshotsLikeDaemon(t *testing.T) {
	m := mock_client.New(t)
	x := t.Context()
	assert.True(
		t,
		validation.Is(
			m.PutSystem(x, "../papa", client.SystemSnapshot{Release: "trixie"}),
		),
	)
	assert.True(
		t,
		validation.Is(m.PutSystem(x, "papa", client.SystemSnapshot{})),
	)
	assert.True(
		t,
		validation.Is(
			m.PutSystem(x, "papa", client.SystemSnapshot{Release: "trixie"}),
		),
	)
	assert.True(
		t,
		validation.Is(
			m.PutSystem(
				x,
				"papa",
				client.SystemSnapshot{
					Release:  "trixie",
					Packages: []*client.InstalledPackage{{Name: "romeo"}},
				},
			),
		),
	)
	assert.FatalOnError(
		t,
		m.PutSystem(
			x,
			"papa",
			client.SystemSnapshot{
				Release: "trixie",
				Packages: []*client.InstalledPackage{
					{
						Name:          "romeo",
						Version:       "1.0-1",
						Source:        "romeo",
						SourceVersion: "1.0-1",
					},
				},
			},
		),
	)
	stored, okay := m.Snapshot("papa")
	assert.True(t, okay)
	assert.String(t, "trixie", stored.Release)
	systems, e := m.ListSystems(x)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, systems)
	assert.False(t, systems[0].Clean)
	assert.FatalOnError(t, m.DeleteSystem(x, "papa"))
	_, f := m.GetSystemFindings(x, "papa")
	assert.True(t, not_found.Is(f))
	assert.True(t, not_found.Is(m.DeleteSystem(x, "papa")))
}

func TestMockAnswersSeededDetailAndGroups(t *testing.T) {
	m := mock_client.New(t)
	m.SeedSystemDetail(
		&client.SystemDetail{
			Summary: client.SystemSummary{Name: "quebec", Clean: true},
		},
	)
	m.SeedFindingGroup(&client.FindingGroup{Vulnerability: "CVE-2099-55555"})
	d, e := m.GetSystemFindings(t.Context(), "quebec")
	assert.FatalOnError(t, e)
	assert.True(t, d.Summary.Clean)
	groups, f := m.ListFindings(t.Context())
	assert.FatalOnError(t, f)
	assert.Count(t, 1, groups)
}
