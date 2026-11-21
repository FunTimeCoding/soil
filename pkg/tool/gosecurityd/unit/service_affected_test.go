package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
)

func TestAffectedListsSystemsByFixKind(t *testing.T) {
	s := security_tester.New(t).UpgradeFleet().System()
	a, e := s.Affected(vulnerabilityConstant.FixtureIdentifier)
	assert.FatalOnError(t, e)
	assert.Integer(t, 4, a.Count.Systems)
	assert.Integer(t, 2, a.Count.Upgrade)
	assert.Integer(t, 1, a.Count.Pending)
	assert.Integer(t, 0, a.Count.Paid)
	assert.Integer(t, 1, a.Count.Open)
	names := make([]string, 0, len(a.Rows))

	for _, r := range a.Rows {
		names = append(names, r.System)
	}

	assert.Strings(t, []string{"bravo", "charlie", "delta", "alfa"}, names)
	assert.String(t, "debian fix not in candidate", a.Rows[2].Kind)
	assert.String(t, "8.6.1-1", a.Rows[2].Fixed)
	assert.String(t, "no fix yet", a.Rows[3].Kind)
	assert.String(t, "romeo", a.Rows[3].Source)
	none, f := s.Affected("CVE-2099-99999")
	assert.FatalOnError(t, f)
	assert.Count(t, 0, none.Rows)
	_, g := s.Affected("bogus")
	assert.True(t, validation.Is(g))
}

func TestHitsCountSystemsOncePerVulnerability(t *testing.T) {
	s := security_tester.New(t).UpgradeFleet().System()
	hits, e := s.Hits(
		[]string{
			vulnerabilityConstant.FixtureIdentifier,
			"CVE-2099-20003",
			"CVE-2099-99999",
		},
	)
	assert.FatalOnError(t, e)
	known := hits[vulnerabilityConstant.FixtureIdentifier]
	assert.Integer(t, 4, known.Systems)
	assert.Integer(t, 1, known.Open)
	assert.Integer(t, 0, hits["CVE-2099-20003"].Systems)
	assert.Integer(t, 0, hits["CVE-2099-99999"].Systems)
}
