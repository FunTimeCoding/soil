package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
)

func TestOverviewCountsAndClassifiesOpenFlags(t *testing.T) {
	o := security_tester.New(t)
	o.NationalDatabase.SeedVulnerability(
		vulnerabilityConstant.FixtureIdentifier,
		"CRITICAL",
		10.0,
	)
	o.NationalDatabase.SeedVulnerability("CVE-2099-22222", "HIGH", 8.1)
	o.SecurityTracker.SeedKnown()
	o.SecurityTracker.SeedStatus(
		"CVE-2099-22222",
		"romeo",
		"trixie",
		vulnerabilityConstant.DebianOpen,
		"",
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	f := o.Flag()

	for _, name := range []string{
		vulnerabilityConstant.FixtureIdentifier,
		"CVE-2099-22222",
		"CVE-2099-33333",
	} {
		_, e := f.Flag(t.Context(), name, "")
		assert.FatalOnError(t, e)
	}

	r, e := o.Overview().Overview("")
	assert.FatalOnError(t, e)
	assert.Integer(t, 3, r.Open)
	assert.Integer(t, 1, r.Critical)
	assert.Integer(t, 1, r.DebianFix)
	assert.Integer(t, 1, r.LocalAction)
	assert.Count(t, 3, r.Rows)
	assert.String(t, "CVE-2099-00001", r.Rows[0].Flag.Subject)
	assert.String(t, "debian fix", r.Rows[0].Fix)
	assert.String(t, "CVE-2099-22222", r.Rows[1].Flag.Subject)
	assert.String(t, "local action", r.Rows[1].Fix)
	assert.String(t, "no tracker data", r.Rows[2].Fix)
}

func TestOverviewCountsOnlyOpenFlags(t *testing.T) {
	o := security_tester.New(t)
	f := o.Flag()
	_, e := f.Flag(t.Context(), vulnerabilityConstant.FixtureIdentifier, "")
	assert.FatalOnError(t, e)
	_, g := f.Mitigate(vulnerabilityConstant.FixtureIdentifier, "")
	assert.FatalOnError(t, g)
	r, h := o.Overview().Overview("all")
	assert.FatalOnError(t, h)
	assert.Integer(t, 0, r.Open)
	assert.Count(t, 1, r.Rows)
}
