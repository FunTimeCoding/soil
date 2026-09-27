//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"testing"
)

func TestCollectionFacets(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"a.md",
			"# Alfa\n\nFirst.\n",
			map[string][]string{
				constant.FixtureAuthorKey: {"alice"},
				constant.FixtureTagKey:    {"design"},
			},
		),
	)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"b.md",
			"# Bravo\n\nSecond.\n",
			map[string][]string{
				constant.FixtureAuthorKey: {"alice"},
				constant.FixtureTagKey:    {constant.FixtureBuildValue},
			},
		),
	)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"c.md",
			"# Charlie\n\nThird.\n",
			map[string][]string{
				constant.FixtureAuthorKey: {"bob"},
				constant.FixtureTagKey:    {constant.FixtureBuildValue},
			},
		),
	)
	facets := s.Service.CollectionFacets("notes", nil)
	authorFacet := findFacet(facets, constant.FixtureAuthorKey)
	assert.NotNil(t, authorFacet)
	assert.Integer(t, 2, authorFacet.Distinct)
	assert.Integer(t, 2, authorFacet.Values["alice"])
	assert.Integer(t, 1, authorFacet.Values["bob"])
	tagFacet := findFacet(facets, constant.FixtureTagKey)
	assert.NotNil(t, tagFacet)
	assert.Integer(t, 2, tagFacet.Distinct)
	assert.Integer(t, 2, tagFacet.Values[constant.FixtureBuildValue])
	assert.Integer(t, 1, tagFacet.Values["design"])
}

func TestCollectionFacetsForKey(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"a.md",
			"# Alfa\n\nFirst.\n",
			map[string][]string{
				constant.FixtureAuthorKey: {"alice"},
				constant.FixtureTagKey:    {"design"},
			},
		),
	)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"b.md",
			"# Bravo\n\nSecond.\n",
			map[string][]string{
				constant.FixtureAuthorKey: {"bob"},
				constant.FixtureTagKey:    {constant.FixtureBuildValue},
			},
		),
	)
	facets := s.Service.CollectionFacetsForKey(
		"notes",
		constant.FixtureAuthorKey,
	)
	assert.Count(t, 1, facets)
	assert.String(t, "author", facets[0].Key)
	assert.Integer(t, 2, facets[0].Distinct)
	assert.Integer(t, 1, facets[0].Values["alice"])
	assert.Integer(t, 1, facets[0].Values["bob"])
}

func TestEmbedCreatesEmbeddings(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	result, e := s.Service.Embed()
	assert.FatalOnError(t, e)
	assert.Greater(t, 0, result.Documents)
	assert.Greater(t, 0, result.Chunks)
	status := s.Service.MustStatus()
	assert.Greater(t, 0, status.TotalEmbeddings)
	assert.Integer(t, 0, status.PendingEmbeddings)
}

func TestEmbedIdempotent(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	first, e := s.Service.Embed()
	assert.FatalOnError(t, e)
	second, e := s.Service.Embed()
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, second.Documents)
	assert.Integer(t, 0, second.Chunks)
	status := s.Service.MustStatus()
	assert.Integer(t, first.Chunks, status.TotalEmbeddings)
}

func TestEmbedAfterPush(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"embed-me.md",
			"# Embed Test\n\nThis document should get embedded.\n",
			nil,
		),
	)
	status := s.Service.MustStatus()
	assert.Greater(t, 0, status.TotalEmbeddings)
}

func TestGetDocument(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	document, similar, e := s.Service.GetDocument("test/alfa.md")
	assert.FatalOnError(t, e)
	assert.NotNil(t, document)
	assert.String(t, "Search Pipeline", document.Title)
	assert.Count(t, 0, similar)
}

func TestGetDocumentNotFoundWithSuggestions(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	document, similar, e := s.Service.GetDocument("test/alfaa.md")
	assert.FatalOnError(t, e)
	assert.Nil(t, document)
	assert.Greater(t, 0, len(similar))
}

func TestGetDocumentNotFoundNoSuggestions(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	document, similar, e := s.Service.GetDocument(
		"test/completely-unrelated.md",
	)
	assert.FatalOnError(t, e)
	assert.Nil(t, document)
	assert.Count(t, 0, similar)
}

func TestGetDocumentMalformedReference(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	document, similar, e := s.Service.GetDocument("completely-unrelated")
	assert.True(t, validation.Is(e))
	assert.Nil(t, document)
	assert.Count(t, 0, similar)
}

func TestIndexCollections(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	results, e := s.Service.IndexCollections("")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, results)
	assert.String(t, "test", results[0].Collection)
	assert.Integer(t, 0, results[0].Indexed)
	assert.Integer(t, 5, results[0].Unchanged)
}

func TestIndexCollectionsFilter(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	results, e := s.Service.IndexCollections("nonexistent")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, results)
}

func TestPushDocumentAppearsInList(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"first.md",
			"# First Note\n\nPushed directly via API.\n",
			nil,
		),
	)
	outcome, e := s.Service.ListDocuments("notes", nil, 0, 0, false)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, outcome.Results)
	assert.String(t, "First Note", outcome.Results[0].Title)
	assert.String(t, "qmd://notes/first.md", outcome.Results[0].VirtualPath)
}

func TestPushDocumentAppearsInSearch(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"searchable.md",
			"# Searchable\n\nThis document contains a unique keyword: platypus.\n",
			nil,
		),
	)
	o := search_option.New("platypus", 10)
	o.Mode = "keyword"
	result := s.Service.Search(o)
	assert.Count(t, 1, result.Results)
	assert.String(t, "Searchable", result.Results[0].Title)
}

func TestPushDocumentDedup(t *testing.T) {
	s := service_tester.New(t)
	body := "# Identical\n\nSame content pushed twice.\n"
	assert.FatalOnError(
		t,
		s.Service.PushDocument("notes", "first.md", body, nil),
	)
	assert.FatalOnError(
		t,
		s.Service.PushDocument("notes", "first.md", body, nil),
	)
	outcome, e := s.Service.ListDocuments("notes", nil, 0, 0, false)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, outcome.Results)
	status := s.Service.MustStatus()
	assert.Integer(t, 1, status.TotalDocuments)
}

func TestPushDocumentUpdate(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"evolving.md",
			"# Original\n\nFirst version.\n",
			nil,
		),
	)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"evolving.md",
			"# Updated\n\nSecond version with revised content.\n",
			nil,
		),
	)
	d, _, e := s.Service.GetDocument("notes/evolving.md")
	assert.FatalOnError(t, e)
	assert.NotNil(t, d)
	assert.String(t, "Updated", d.Title)
	assert.StringContains(t, "revised content", d.Body)
}

func TestPushDocumentCreatesCollection(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"dynamic",
			"auto.md",
			"# Auto\n\nCollection created on first push.\n",
			nil,
		),
	)
	status := s.Service.MustStatus()
	found := false

	for _, c := range status.Collections {
		if c.Name == "dynamic" {
			found = true
			assert.Integer(t, 1, c.DocumentCount)
		}
	}

	assert.True(t, found)
}

func TestSearchKeyword(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	o := search_option.New("hybrid search pipeline", 10)
	o.Mode = "keyword"
	outcome := s.Service.Search(o)
	assert.Greater(t, 0, len(outcome.Results))
	assert.String(t, "Search Pipeline", outcome.Results[0].Title)
}

func TestSearchHybrid(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	_, e := s.Service.Embed()
	assert.FatalOnError(t, e)
	outcome := s.Service.Search(search_option.New("semantic meaning", 10))
	assert.Greater(t, 0, len(outcome.Results))
}

func TestSearchCollectionFilter(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	o := search_option.New("chunking", 10)
	o.Collection = "test"
	o.Mode = "keyword"
	outcome := s.Service.Search(o)
	assert.Greater(t, 0, len(outcome.Results))
	assert.String(t, "test", outcome.Results[0].Collection)
}

func TestSearchDegradedWithoutEmbeddings(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	outcome := s.Service.Search(search_option.New("chunking strategy", 10))
	assert.True(t, outcome.Degraded)
	assert.Greater(t, 0, len(outcome.Results))
}
