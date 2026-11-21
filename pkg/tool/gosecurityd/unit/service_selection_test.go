package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	"testing"
)

func TestBoardSelectsByLabelAndCountsFacets(t *testing.T) {
	s := security_tester.New(t).LabeledFleet().System()
	b, e := s.Board(system_selection.New().WithLabel("exposure", "internet"))
	assert.FatalOnError(t, e)
	names := make([]string, 0, len(b.Rows))

	for _, v := range b.Rows {
		names = append(names, v.Name)
	}

	assert.Strings(t, []string{"alfa", "bravo"}, names)
	assert.Integer(t, 2, b.Total)
	assert.Integer(t, 2, b.Upgradable)
	assert.Integer(t, 4, b.LabelCounts["exposure"][""])
	assert.Integer(t, 2, b.LabelCounts["exposure"]["internet"])
	assert.Integer(t, 2, b.LabelCounts["exposure"]["internal"])
	assert.Integer(t, 2, b.LabelCounts["tenant"][""])
	assert.Integer(t, 1, b.LabelCounts["tenant"]["one"])
	assert.Integer(t, 1, b.LabelCounts["tenant"]["two"])
	assert.Integer(t, 2, b.CoverageCounts[""])
	assert.Integer(t, 2, b.CoverageCounts[constant.ShowSupported])
	assert.Integer(t, 2, b.ScoreCounts[0])
	both, f := s.Board(
		system_selection.New().
			WithLabel("exposure", "internal").
			WithLabel("tenant", "one"),
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, both.Rows)
	assert.String(t, "charlie", both.Rows[0].Name)
	assert.Integer(t, 1, both.LabelCounts["exposure"]["internet"])
}

func TestBoardSearchesNamesIgnoringCase(t *testing.T) {
	s := security_tester.New(t).LabeledFleet().System()
	b, e := s.Board(system_selection.New().WithSearch("RAV"))
	assert.FatalOnError(t, e)
	assert.Count(t, 1, b.Rows)
	assert.String(t, "bravo", b.Rows[0].Name)
	assert.Integer(t, 1, b.Total)
	assert.Integer(t, 1, b.LabelCounts["exposure"]["internet"])
}

func TestSelectionBuildsLinksWithoutChangingItself(t *testing.T) {
	s := system_selection.New().
		WithCoverage(constant.ShowSupported).
		WithMinimum(9).
		WithLabel("tenant", "one").
		WithLabel("exposure", "internet").
		WithSearch("al")
	assert.String(
		t,
		"label=exposure%3Ainternet&label=tenant%3Aone&score=9&search=al&show=supported",
		s.Query().Encode(),
	)
	changed := s.WithLabel("tenant", "").WithCoverage("")
	assert.String(
		t,
		"label=exposure%3Ainternet&score=9&search=al",
		changed.Query().Encode(),
	)
	assert.String(
		t,
		"label=exposure%3Ainternet&label=tenant%3Aone&score=9&search=al&show=supported",
		s.Query().Encode(),
	)
	assert.True(
		t,
		validation.Is(system_selection.New().WithCoverage("x").Check()),
	)
	assert.True(
		t,
		validation.Is(system_selection.New().WithMinimum(11).Check()),
	)
	assert.FatalOnError(t, s.Check())
}
