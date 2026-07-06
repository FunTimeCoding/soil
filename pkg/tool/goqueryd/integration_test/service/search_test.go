//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"testing"
)

func TestSearchKeyword(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	o := search_option.New("hybrid search pipeline", 10)
	o.Mode = "keyword"
	outcome := s.Service.Search(o)
	assert.Greater(t, 0, float64(len(outcome.Results)))
	assert.String(t, "Search Pipeline", outcome.Results[0].Title)
}

func TestSearchHybrid(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	_, e := s.Service.Embed()
	assert.FatalOnError(t, e)
	outcome := s.Service.Search(search_option.New("semantic meaning", 10))
	assert.Greater(t, 0, float64(len(outcome.Results)))
}

func TestSearchCollectionFilter(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	o := search_option.New("chunking", 10)
	o.Collection = "test"
	o.Mode = "keyword"
	outcome := s.Service.Search(o)
	assert.Greater(t, 0, float64(len(outcome.Results)))
	assert.String(t, "test", outcome.Results[0].Collection)
}

func TestSearchDegradedWithoutEmbeddings(t *testing.T) {
	s := service_tester.New(t)
	s.IndexFixtures()
	outcome := s.Service.Search(search_option.New("chunking strategy", 10))
	assert.True(t, outcome.Degraded)
	assert.Greater(t, 0, float64(len(outcome.Results)))
}
