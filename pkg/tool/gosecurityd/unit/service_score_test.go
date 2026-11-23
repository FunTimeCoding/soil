package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	vulnerabilityConstant "github.com/funtimecoding/soil/pkg/vulnerability/constant"
	"testing"
	"time"
)

func TestScoreFillPagesThroughCatalogue(t *testing.T) {
	o := security_tester.New(t).SeedScores()
	o.NationalDatabase.PageSize(2)
	s := o.Score()
	assert.Count(t, 0, s.Sync(t.Context()))
	assert.Integers(t, []int{0, 2, 4}, o.NationalDatabase.Starts())
	status := s.Status()
	assert.True(t, status.Filled)
	assert.Integer(t, 5, status.Mirrored)
	assert.Integer(t, 5, status.Total)
	c, okay := s.Score("CVE-2099-90001")
	assert.True(t, okay)
	assert.Float(t, 9.8, c.Score)
}

func TestScoreFillResumesWhereItStopped(t *testing.T) {
	o := security_tester.New(t).SeedScores()
	o.NationalDatabase.PageSize(2)
	o.NationalDatabase.FailAfter(1)
	s := o.Score()
	assert.Count(t, 1, s.Sync(t.Context()))
	assert.False(t, s.Status().Filled)
	assert.Integer(t, 2, s.Status().Start)
	o.NationalDatabase.FailAfter(0)
	assert.Count(t, 0, s.Sync(t.Context()))
	assert.Integers(t, []int{0, 2, 2, 4}, o.NationalDatabase.Starts())
	assert.True(t, s.Status().Filled)
	assert.Integer(t, 5, s.Status().Mirrored)
}

func TestScoresSurviveRestart(t *testing.T) {
	o := security_tester.New(t).SeedScores()
	assert.Count(t, 0, o.Score().Sync(t.Context()))
	o.NationalDatabase.Fail(errors.New("victor"))
	restarted := o.FreshScore()
	assert.Count(t, 1, restarted.Sync(t.Context()))
	assert.True(t, restarted.Status().Filled)
	assert.Integer(t, 5, restarted.Status().Mirrored)
	c, okay := restarted.Score(vulnerabilityConstant.FixtureIdentifier)
	assert.True(t, okay)
	assert.Float(t, 10.0, c.Score)
}

func TestScoreUpdateFetchesChangesInWindows(t *testing.T) {
	o := security_tester.New(t).SeedScores()
	o.Later = -300 * 24 * time.Hour
	s := o.Score()
	assert.Count(t, 0, s.Sync(t.Context()))
	o.Later = 0
	o.NationalDatabase.SeedScore(
		"CVE-2099-90009",
		9.1,
		"CRITICAL",
		time.Now().Add(-time.Hour),
	)
	assert.Count(t, 0, s.Sync(t.Context()))
	assert.Count(t, 3, o.NationalDatabase.Windows())
	c, okay := s.Score("CVE-2099-90009")
	assert.True(t, okay)
	assert.Float(t, 9.1, c.Score)
	assert.Time(t, time.Now(), s.Status().Synced)
}

func TestFindingsCarryScoresAndFollowChanges(t *testing.T) {
	o := security_tester.New(t).UpgradeFleet()
	assert.Count(t, 0, o.Score().Sync(t.Context()))
	s := o.System()
	before, e := s.Detail("bravo")
	assert.FatalOnError(t, e)
	assert.Float(t, 0, before.Findings[0].Score)
	o.NationalDatabase.SeedScore(
		vulnerabilityConstant.FixtureIdentifier,
		10.0,
		"CRITICAL",
		time.Now(),
	)
	assert.Count(t, 0, o.Score().Sync(t.Context()))
	after, f := s.Detail("bravo")
	assert.FatalOnError(t, f)
	scored := map[string]float64{}

	for _, v := range after.Findings {
		scored[v.Vulnerability] = v.Score
	}

	assert.Float(t, 10.0, scored[vulnerabilityConstant.FixtureIdentifier])
	u, g := s.Upgrades(system_selection.New(), "", false)
	assert.FatalOnError(t, g)
	assert.Float(t, 10.0, u.Groups[0].TopScore)
	a, h := s.Affected(vulnerabilityConstant.FixtureIdentifier)
	assert.FatalOnError(t, h)
	assert.Float(t, 10.0, a.Score)
	assert.String(t, "CRITICAL", a.Severity)
	p, i := s.Packages("bravo", 0)
	assert.FatalOnError(t, i)
	assert.Float(t, 10.0, p.Upgrade[0].TopScore)
}
