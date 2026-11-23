package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/security_tester"
	"testing"
)

func TestBoardSelectsByLabel(t *testing.T) {
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
	both, f := s.Board(
		system_selection.New().
			WithLabel("exposure", "internal").
			WithLabel("tenant", "one"),
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, both.Rows)
	assert.String(t, "charlie", both.Rows[0].Name)
}

func TestFacetsCountEachChipWithOtherFiltersApplied(t *testing.T) {
	s := security_tester.New(t).LabeledFleet().System()
	c, e := s.Facets(system_selection.New().WithLabel("exposure", "internet"))
	assert.FatalOnError(t, e)
	assert.Integer(t, 4, c.Labels["exposure"][""])
	assert.Integer(t, 2, c.Labels["exposure"]["internet"])
	assert.Integer(t, 2, c.Labels["exposure"]["internal"])
	assert.Integer(t, 2, c.Labels["tenant"][""])
	assert.Integer(t, 1, c.Labels["tenant"]["one"])
	assert.Integer(t, 1, c.Labels["tenant"]["two"])
	assert.Integer(t, 2, c.Coverage[""])
	assert.Integer(t, 2, c.Coverage[constant.ShowSupported])
	assert.Integer(t, 2, c.Score[0])
	both, f := s.Facets(
		system_selection.New().
			WithLabel("exposure", "internal").
			WithLabel("tenant", "one"),
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, both.Labels["exposure"]["internet"])
	_, g := s.Facets(system_selection.New().WithCoverage("bogus"))
	assert.True(t, validation.Is(g))
}

func TestBoardSearchesNamesIgnoringCase(t *testing.T) {
	s := security_tester.New(t).LabeledFleet().System()
	b, e := s.Board(system_selection.New().WithSearch("RAV"))
	assert.FatalOnError(t, e)
	assert.Count(t, 1, b.Rows)
	assert.String(t, "bravo", b.Rows[0].Name)
	assert.Integer(t, 1, b.Total)
	c, f := s.Facets(system_selection.New().WithSearch("RAV"))
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, c.Labels["exposure"]["internet"])
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
		s.Query(0).Encode(),
	)
	changed := s.WithLabel("tenant", "").WithCoverage("")
	assert.String(
		t,
		"label=exposure%3Ainternet&score=9&search=al",
		changed.Query(0).Encode(),
	)
	assert.String(
		t,
		"label=exposure%3Ainternet&label=tenant%3Aone&score=9&search=al&show=supported",
		s.Query(0).Encode(),
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

func TestSelectionLinksCarryScoreOnlyAwayFromDefault(t *testing.T) {
	critical := system_selection.New().WithMinimum(9)
	assert.String(t, "", critical.Query(9).Encode())
	assert.String(t, "score=9", critical.Query(0).Encode())
	assert.String(t, "score=0", critical.WithMinimum(0).Query(9).Encode())
}

func TestSelectionReadsLabelText(t *testing.T) {
	s, e := system_selection.New().WithLabelText("group:red")
	assert.FatalOnError(t, e)
	assert.String(t, "red", s.Labels["group"])

	for _, v := range []string{"group", "group:", ":red", ""} {
		_, f := system_selection.New().WithLabelText(v)
		assert.True(t, validation.Is(f))
	}
}
