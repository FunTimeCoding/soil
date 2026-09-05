package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"testing"
)

func taggedStore(t *testing.T) *store.Store {
	t.Helper()
	s := openTestStore(t)
	directory := t.TempDir()
	writeFixture(t, directory, "alfa.md", "# Alfa\n\nKeyword quasar.\n")
	writeFixture(t, directory, "bravo.md", "# Bravo\n\nKeyword quasar.\n")
	writeFixture(t, directory, "charlie.md", "# Charlie\n\nKeyword quasar.\n")
	s.AddCollection("test", directory, constant.DefaultGlob)
	s.Index("test")
	s.SetMetadata(
		"test",
		"alfa.md",
		map[string][]string{
			constant.FixtureTagKey:    {constant.FixtureBuildValue, "groom"},
			constant.FixtureAuthorKey: {"one"},
		},
	)
	s.SetMetadata(
		"test",
		"bravo.md",
		map[string][]string{
			constant.FixtureTagKey:    {constant.FixtureBuildValue},
			constant.FixtureAuthorKey: {"two"},
		},
	)
	s.SetMetadata(
		"test",
		"charlie.md",
		map[string][]string{
			constant.FixtureTagKey:    {"groom"},
			constant.FixtureAuthorKey: {"one"},
		},
	)

	return s
}

func tagFilter(value string) map[string]string {
	return map[string]string{constant.FixtureTagKey: value}
}

func TestDocumentIsFoundByEitherOfItsValues(t *testing.T) {
	s := taggedStore(t)
	defer s.Close()
	build := s.MustSearchKeyword(
		"quasar",
		10,
		"",
		false,
		tagFilter(constant.FixtureBuildValue),
	)
	assert.Count(t, 2, build)
	groom := s.MustSearchKeyword("quasar", 10, "", false, tagFilter("groom"))
	assert.Count(t, 2, groom)
}

func TestAbsentValueMatchesNothing(t *testing.T) {
	s := taggedStore(t)
	defer s.Close()
	assert.Count(
		t,
		0,
		s.MustSearchKeyword("quasar", 10, "", false, tagFilter("absent")),
	)
}

func TestTwoFiltersNarrowTogether(t *testing.T) {
	s := taggedStore(t)
	defer s.Close()
	results := s.MustSearchKeyword(
		"quasar",
		10,
		"",
		false,
		map[string]string{
			constant.FixtureTagKey:    constant.FixtureBuildValue,
			constant.FixtureAuthorKey: "one",
		},
	)
	assert.Count(t, 1, results)
	assert.String(t, "Alfa", results[0].Title)
}

func TestReplaceRemovesDroppedValues(t *testing.T) {
	s := taggedStore(t)
	defer s.Close()
	s.SetMetadata(
		"test",
		"alfa.md",
		map[string][]string{
			constant.FixtureTagKey: {constant.FixtureBuildValue},
		},
	)
	assert.Count(
		t,
		1,
		s.MustSearchKeyword("quasar", 10, "", false, tagFilter("groom")),
	)
}

func TestScalarKeyStaysScalar(t *testing.T) {
	s := taggedStore(t)
	defer s.Close()
	filter := map[string]string{constant.FixtureAuthorKey: "one"}
	results := s.EnrichResults(
		s.MustSearchKeyword("quasar", 10, "", false, filter),
		filter,
	)
	assert.Count(t, 2, results)
	assert.Count(t, 1, results[0].Metadata[constant.FixtureAuthorKey])
}

func TestResultCarriesEveryValue(t *testing.T) {
	s := taggedStore(t)
	defer s.Close()
	two := map[string]string{constant.FixtureAuthorKey: "two"}
	results := s.EnrichResults(
		s.MustSearchKeyword("quasar", 10, "", false, two),
		two,
	)
	assert.Count(t, 1, results)
	assert.Count(t, 1, results[0].Metadata[constant.FixtureTagKey])
	one := map[string]string{constant.FixtureAuthorKey: "one"}
	alfa := s.EnrichResults(
		s.MustSearchKeyword("quasar", 10, "", false, one),
		one,
	)
	assert.Count(t, 2, alfa[0].Metadata[constant.FixtureTagKey])
}

func TestCollectionFacetsCountEachValue(t *testing.T) {
	s := taggedStore(t)
	defer s.Close()
	facets := s.CollectionFacets("test", nil, 20, constant.FixtureTagKey)
	tag := findFacet(facets, constant.FixtureTagKey)
	assert.NotNil(t, tag)
	assert.Integer(t, 2, tag.Distinct)
	assert.Integer(t, 2, tag.Values[constant.FixtureBuildValue])
	assert.Integer(t, 2, tag.Values["groom"])
}

func TestPanelQueryShape(t *testing.T) {
	s := taggedStore(t)
	defer s.Close()
	filter := tagFilter(constant.FixtureBuildValue)
	found := s.MustSearchKeyword("quasar", 10, "", false, filter)
	unread := s.EnrichResults(
		store.ExcludePaths(found, []string{"alfa.md"}),
		filter,
	)
	assert.Count(t, 1, unread)
	assert.String(t, "Bravo", unread[0].Title)
	assert.Count(t, 1, unread[0].Metadata[constant.FixtureTagKey])
}

func TestRepeatedValueStoresOnce(t *testing.T) {
	s := taggedStore(t)
	defer s.Close()
	s.SetMetadata(
		"test",
		"alfa.md",
		map[string][]string{
			constant.FixtureTagKey: {
				constant.FixtureBuildValue,
				constant.FixtureBuildValue,
			},
		},
	)
	filter := tagFilter(constant.FixtureBuildValue)
	results := s.EnrichResults(
		s.MustSearchKeyword("quasar", 10, "", false, filter),
		filter,
	)
	assert.Count(t, 2, results)
	assert.Count(t, 1, results[0].Metadata[constant.FixtureTagKey])
}
