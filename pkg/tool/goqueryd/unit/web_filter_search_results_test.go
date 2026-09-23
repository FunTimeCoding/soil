package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web"
	"testing"
)

func TestFilterSearchResultsWithoutFilterKeepsEverything(t *testing.T) {
	results, facets := web.FilterSearchResults(
		store_tester.FilterOutcome(),
		nil,
	)
	assert.Count(t, 3, results)
	assert.True(t, len(facets) > 0)
}

func TestFilterSearchResultsMatchesOnOneValue(t *testing.T) {
	results, _ := web.FilterSearchResults(
		store_tester.FilterOutcome(),
		map[string]string{constant.FixtureTagKey: constant.FixtureBuildValue},
	)
	assert.Count(t, 2, results)
	assert.String(t, "alfa", results[0].Title)
	assert.String(t, "charlie", results[1].Title)
}

func TestFilterSearchResultsExcludesNonMatching(t *testing.T) {
	results, _ := web.FilterSearchResults(
		store_tester.FilterOutcome(),
		map[string]string{constant.FixtureTagKey: "absent"},
	)
	assert.Count(t, 0, results)
}

func TestFilterSearchResultsFacetsFollowTheFilter(t *testing.T) {
	_, facets := web.FilterSearchResults(
		store_tester.FilterOutcome(),
		map[string]string{constant.FixtureAuthorKey: "one"},
	)
	author := store_tester.FindFacet(facets, constant.FixtureAuthorKey)
	assert.NotNil(t, author)
	assert.Integer(t, 1, author.Distinct)
	assert.Integer(t, 2, author.Values["one"])
	assert.Nil(t, store_tester.FindFacet(facets, constant.FixtureTagKey))
}

func TestFilterSearchResultsSuppressesUniquePerResultKeys(t *testing.T) {
	outcome := store.NewSearchOutcome(
		[]store.SearchResult{
			{
				Title:    "alfa",
				Metadata: map[string][]string{constant.Path: {"one"}},
			},
			{
				Title:    "bravo",
				Metadata: map[string][]string{constant.Path: {"two"}},
			},
		},
	)
	_, facets := web.FilterSearchResults(outcome, nil)
	assert.Nil(t, store_tester.FindFacet(facets, constant.Path))
}
