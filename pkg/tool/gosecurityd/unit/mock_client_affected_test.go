package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"testing"
)

func TestMockAnswersAffectedSystemsLikeDaemon(t *testing.T) {
	m := mock_client.New(t)
	x := t.Context()
	m.SeedAffectedSystems(
		&client.AffectedSystems{
			Identifier: "CVE-2099-50001",
			Count:      client.AffectedCount{Systems: 1, Upgrade: 1},
			Systems:    []client.AffectedSystem{{System: "india"}},
		},
	)
	seeded, e := m.GetAffectedSystems(x, "CVE-2099-50001")
	assert.FatalOnError(t, e)
	assert.String(t, "india", seeded.Systems[0].System)
	empty, f := m.GetAffectedSystems(x, "CVE-2099-50002")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, empty.Systems)
	_, g := m.GetAffectedSystems(x, "bogus")
	assert.True(t, validation.Is(g))
}
