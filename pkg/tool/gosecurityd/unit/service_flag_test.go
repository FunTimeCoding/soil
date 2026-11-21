package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
)

func TestFlagKeepsDatabaseSnapshot(t *testing.T) {
	o := security_tester.New(t)
	o.NationalDatabase.SeedVulnerability(
		vulnerabilityConstant.FixtureIdentifier,
		"CRITICAL",
		10.0,
	)
	f, e := o.Flag().Flag(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		"alfa",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "CRITICAL", f.Severity)
	assert.Float(t, 10.0, f.Score)
	assert.String(t, "alfa", f.Note)
	flags, g := o.Flag().Flags(false)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, flags)
	assert.String(t, "CVE-2099-00001", flags[0].Subject)
}

func TestFlagWithoutDatabaseRecordStillFlags(t *testing.T) {
	o := security_tester.New(t)
	f, e := o.Flag().Flag(t.Context(), "CVE-2099-99999", "")
	assert.FatalOnError(t, e)
	assert.String(t, "", f.Severity)
}

func TestFlagRefusesWhenDatabaseFails(t *testing.T) {
	o := security_tester.New(t)
	o.NationalDatabase.Fail(errors.New("bravo"))
	_, e := o.Flag().Flag(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		"",
	)
	assert.Error(t, e)
	flags, f := o.Flag().Flags(true)
	assert.FatalOnError(t, f)
	assert.Count(t, 0, flags)
}

func TestFlagRemediationItemWithoutDatabase(t *testing.T) {
	o := security_tester.New(t)
	o.NationalDatabase.Fail(errors.New("bravo"))
	f, e := o.Flag().Flag(t.Context(), "upgrade:bookworm:victor", "hold")
	assert.FatalOnError(t, e)
	assert.String(t, "", f.Severity)
	m, g := o.Flag().Mitigate("upgrade:bookworm:victor", "rolled out")
	assert.FatalOnError(t, g)
	assert.True(t, m.Mitigated)
	_, h := o.Flag().Flag(t.Context(), "upgrade:bookworm", "")
	assert.True(t, validation.Is(h))
}

func TestFlaggingOpenFlagAgainIsConflict(t *testing.T) {
	o := security_tester.New(t)
	_, e := o.Flag().Flag(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		"",
	)
	assert.FatalOnError(t, e)
	_, f := o.Flag().Flag(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		"",
	)
	assert.True(t, conflict.Is(f))
}

func TestMitigatedFlagLeavesOpenList(t *testing.T) {
	o := security_tester.New(t)
	_, e := o.Flag().Flag(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		"",
	)
	assert.FatalOnError(t, e)
	m, f := o.Flag().Mitigate(
		vulnerabilityConstant.FixtureIdentifier,
		"charlie",
	)
	assert.FatalOnError(t, f)
	assert.True(t, m.Mitigated)
	assert.String(t, "charlie", m.Mitigation)
	open, g := o.Flag().Flags(false)
	assert.FatalOnError(t, g)
	assert.Count(t, 0, open)
	all, h := o.Flag().Flags(true)
	assert.FatalOnError(t, h)
	assert.Count(t, 1, all)
}

func TestMitigatingTwiceIsConflict(t *testing.T) {
	o := security_tester.New(t)
	_, e := o.Flag().Flag(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		"",
	)
	assert.FatalOnError(t, e)
	_, f := o.Flag().Mitigate(vulnerabilityConstant.FixtureIdentifier, "")
	assert.FatalOnError(t, f)
	_, g := o.Flag().Mitigate(vulnerabilityConstant.FixtureIdentifier, "")
	assert.True(t, conflict.Is(g))
}

func TestMitigatingUnflaggedIsNotFound(t *testing.T) {
	_, e := security_tester.New(t).Flag().Mitigate(
		vulnerabilityConstant.FixtureIdentifier,
		"",
	)
	assert.True(t, not_found.Is(e))
}

func TestFlaggingMitigatedReopens(t *testing.T) {
	o := security_tester.New(t)
	_, e := o.Flag().Flag(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		"",
	)
	assert.FatalOnError(t, e)
	_, f := o.Flag().Mitigate(vulnerabilityConstant.FixtureIdentifier, "delta")
	assert.FatalOnError(t, f)
	r, g := o.Flag().Flag(
		t.Context(),
		vulnerabilityConstant.FixtureIdentifier,
		"echo",
	)
	assert.FatalOnError(t, g)
	assert.False(t, r.Mitigated)
	assert.String(t, "", r.Mitigation)
	assert.String(t, "echo", r.Note)
	all, h := o.Flag().Flags(true)
	assert.FatalOnError(t, h)
	assert.Count(t, 1, all)
}

func TestFlagRefusesMalformedIdentifier(t *testing.T) {
	o := security_tester.New(t)
	_, e := o.Flag().Flag(t.Context(), "foxtrot", "")
	assert.True(t, validation.Is(e))
	_, f := o.Flag().Mitigate("foxtrot", "")
	assert.True(t, validation.Is(f))
	assert.Integer(t, 0, o.NationalDatabase.Calls())
}
