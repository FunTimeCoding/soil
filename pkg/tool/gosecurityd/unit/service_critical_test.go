package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"math"
	"testing"
	"time"
)

func TestSummariesCountCriticalAndNarrowByScore(t *testing.T) {
	s := security_tester.New(t).ScoredFleet().System()
	all, e := s.Summaries(0)
	assert.FatalOnError(t, e)
	assert.Count(t, 5, all)
	critical := map[string]int{}
	top := map[string]float64{}

	for _, v := range all {
		critical[v.Name] = v.Critical
		top[v.Name] = v.TopScore
	}

	assert.Integer(t, 2, critical["echo"])
	assert.Integer(t, 1, critical["alfa"])
	assert.Integer(t, 0, critical["charlie"])
	assert.Float(t, 7.5, top["charlie"])
	assert.Float(t, 9.8, top["echo"])
	narrowed, f := s.Summaries(9)
	assert.FatalOnError(t, f)
	assert.Count(t, 4, narrowed)

	for _, v := range narrowed {
		assert.True(t, v.Name != "charlie")
	}
}

func TestCriticalCountsVulnerabilityOnceAcrossPackages(t *testing.T) {
	o := security_tester.New(t)

	for _, source := range []string{"tango-utils", "romeo"} {
		o.SecurityTracker.SeedStatus(
			vulnerabilityConstant.FixtureIdentifier,
			source,
			"trixie",
			vulnerabilityConstant.DebianOpen,
			"",
		)
	}

	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	o.NationalDatabase.SeedScore(
		vulnerabilityConstant.FixtureIdentifier,
		9.8,
		"CRITICAL",
		time.Now(),
	)
	assert.Count(t, 0, o.Score().Sync(t.Context()))
	s := o.System()
	assert.FatalOnError(
		t,
		s.Put(
			system.New("sierra", "trixie", 0),
			[]installed.Installed{
				*installed.New(
					"libtango5",
					"8.8.1-1",
					"tango-utils",
					"8.8.1-1",
					"",
				),
				*installed.New("romeo", "1.0-1", "romeo", "1.0-1", ""),
			},
		),
	)
	d, e := s.Detail("sierra")
	assert.FatalOnError(t, e)
	assert.Count(t, 2, d.Findings)
	assert.Integer(t, 1, d.Summary.Critical)
}

func TestBoardNarrowsByScoreAndRanksCriticalWithinPriority(t *testing.T) {
	s := security_tester.New(t).ScoredFleet().System()
	b, e := s.Board(system_selection.New().WithMinimum(9))
	assert.FatalOnError(t, e)
	assert.Integer(t, 4, b.Total)
	assert.Integer(t, 4, b.Critical)
	assert.Count(t, 4, b.Rows)
	c, g := s.Facets(system_selection.New().WithMinimum(9))
	assert.FatalOnError(t, g)
	assert.Integer(t, 5, c.Score[0])
	assert.Integer(t, 4, c.Score[9])
	high, f := s.Board(system_selection.New().WithMinimum(7))
	assert.FatalOnError(t, f)
	names := make([]string, 0, len(high.Rows))

	for _, v := range high.Rows {
		names = append(names, v.Name)
	}

	assert.Strings(
		t,
		[]string{"delta", "echo", "bravo", "charlie", "alfa"},
		names,
	)
}

func TestUpgradesNarrowToCriticalFixes(t *testing.T) {
	s := security_tester.New(t).ScoredFleet().System()
	all, e := s.Upgrades(system_selection.New(), "", false)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, all.Groups)
	r, f := s.Upgrades(system_selection.New().WithMinimum(9), "", false)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, r.Groups)
	assert.String(t, "trixie", r.Groups[0].Release)
	assert.Strings(t, []string{"CVE-2099-20002"}, r.Groups[0].Vulnerabilities)
	assert.Float(t, 9.8, r.Groups[0].TopScore)
	assert.Integer(t, 1, r.Vulnerabilities)
}

func TestUpgradesWithScoreKeepHostsTimesVulnerabilitiesOrder(t *testing.T) {
	o := security_tester.New(t)
	o.SecurityTracker.SeedStatus(
		"CVE-2099-70001",
		"alpha",
		"trixie",
		vulnerabilityConstant.DebianResolved,
		"2.0-1",
	)
	o.SecurityTracker.SeedStatus(
		"CVE-2099-70002",
		"omega",
		"trixie",
		vulnerabilityConstant.DebianResolved,
		"2.0-1",
	)
	assert.Count(t, 0, o.Debian().Refresh(t.Context()))
	s := o.System()

	for _, v := range []string{"kilo", "lima"} {
		packages := []installed.Installed{
			*installed.New("alpha", "1.0-1", "alpha", "1.0-1", "2.0-1"),
		}

		if v == "lima" {
			packages = append(
				packages,
				*installed.New("omega", "1.0-1", "omega", "1.0-1", "2.0-1"),
			)
		}

		assert.FatalOnError(t, s.Put(system.New(v, "trixie", 0), packages))
	}

	o.NationalDatabase.SeedScore("CVE-2099-70001", 9.1, "CRITICAL", time.Now())
	o.NationalDatabase.SeedScore("CVE-2099-70002", 9.9, "CRITICAL", time.Now())
	assert.Count(t, 0, o.Score().Sync(t.Context()))
	plain, e := s.Upgrades(system_selection.New(), "", false)
	assert.FatalOnError(t, e)
	assert.String(t, "alpha", plain.Groups[0].Source)
	ranked, f := s.Upgrades(system_selection.New().WithMinimum(9), "", false)
	assert.FatalOnError(t, f)
	assert.String(t, "alpha", ranked.Groups[0].Source)
	assert.String(t, "omega", ranked.Groups[1].Source)
}

func TestPackagesNarrowToScore(t *testing.T) {
	s := security_tester.New(t).ScoredFleet().System()
	all, e := s.Packages("alfa", 0)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, all.Upgrade[0].Vulnerabilities)
	assert.Count(t, 1, all.Open)
	p, f := s.Packages("alfa", 9)
	assert.FatalOnError(t, f)
	assert.Strings(t, []string{"CVE-2099-20002"}, p.Upgrade[0].Vulnerabilities)
	assert.Count(t, 0, p.Open)
}

func TestGroupsRankByScoreAndCountWorstKind(t *testing.T) {
	s := security_tester.New(t).ScoredFleet().System()
	all, e := s.Groups(0)
	assert.FatalOnError(t, e)
	assert.String(t, "CVE-2099-20002", all[0].Vulnerability)
	known := all[len(all)-1]
	assert.String(t, "CVE-2099-00001", known.Vulnerability)
	assert.Float(t, 7.5, known.Score)
	assert.String(t, "HIGH", known.Severity)
	assert.Integer(t, 5, known.Count.Systems)
	assert.Integer(t, 2, known.Count.Upgrade)
	assert.Integer(t, 2, known.Count.Pending)
	assert.Integer(t, 1, known.Count.Open)
	critical, f := s.Groups(9)
	assert.FatalOnError(t, f)

	for _, g := range critical {
		assert.True(t, g.Score >= 9)
	}

	assert.Count(t, len(all)-1, critical)
	assert.Strings(
		t,
		[]string{"alfa", "bravo", "delta", "echo"},
		critical[0].Systems,
	)
}

func TestScoreOutsideRangeIsValidation(t *testing.T) {
	s := security_tester.New(t).ScoredFleet().System()

	for _, v := range []float64{-1, 10.5, math.NaN()} {
		_, e := s.Summaries(v)
		assert.True(t, validation.Is(e))
		_, f := s.Board(system_selection.New().WithMinimum(v))
		assert.True(t, validation.Is(f))
		_, g := s.Upgrades(system_selection.New().WithMinimum(v), "", false)
		assert.True(t, validation.Is(g))
		_, h := s.Packages("alfa", v)
		assert.True(t, validation.Is(h))
		_, i := s.Groups(v)
		assert.True(t, validation.Is(i))
	}

	_, e := s.Groups(10)
	assert.FatalOnError(t, e)
}
