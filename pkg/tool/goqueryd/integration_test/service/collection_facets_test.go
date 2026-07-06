//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/service_tester"
	"testing"
)

func TestCollectionFacets(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"a.md",
			"# Alpha\n\nFirst.\n",
			map[string]string{
				constant.FixtureAuthorKey: "alice",
				constant.FixtureTagKey:    "design",
			},
		),
	)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"b.md",
			"# Beta\n\nSecond.\n",
			map[string]string{
				constant.FixtureAuthorKey: "alice",
				constant.FixtureTagKey:    constant.FixtureBuildValue,
			},
		),
	)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"c.md",
			"# Gamma\n\nThird.\n",
			map[string]string{
				constant.FixtureAuthorKey: "bob",
				constant.FixtureTagKey:    constant.FixtureBuildValue,
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
			"# Alpha\n\nFirst.\n",
			map[string]string{
				constant.FixtureAuthorKey: "alice",
				constant.FixtureTagKey:    "design",
			},
		),
	)
	assert.FatalOnError(
		t,
		s.Service.PushDocument(
			"notes",
			"b.md",
			"# Beta\n\nSecond.\n",
			map[string]string{
				constant.FixtureAuthorKey: "bob",
				constant.FixtureTagKey:    constant.FixtureBuildValue,
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
