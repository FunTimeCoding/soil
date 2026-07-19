//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"testing"
)

func TestHybridSearchReturnsResults(t *testing.T) {
	s, o := indexedTestStore(t)
	defer s.Close()
	e := embedTestDocuments(s, o)
	assert.FatalOnError(t, e)
	option := search_option.New("search pipeline", 10)
	results, e := s.SearchHybrid(option, o)
	assert.FatalOnError(t, e)
	assert.Greater(t, 0, float64(len(results)))
}

func TestHybridSearchDiffersFromKeyword(t *testing.T) {
	s, o := indexedTestStore(t)
	defer s.Close()
	e := embedTestDocuments(s, o)
	assert.FatalOnError(t, e)
	keyword := s.MustSearchKeyword("context resolution", 5, "", false)
	option := search_option.New("context resolution", 5)
	hybrid, e := s.SearchHybrid(option, o)
	assert.FatalOnError(t, e)

	if len(keyword) > 1 && len(hybrid) > 1 {
		keywordFirst := keyword[0].FilePath
		hybridFirst := hybrid[0].FilePath
		keywordScore := keyword[0].Score
		hybridScore := hybrid[0].Score

		if keywordFirst == hybridFirst {
			assert.True(t, keywordScore != hybridScore)
		}
	}
}

func TestSearchFallsBackToKeywordWithoutEmbeddings(t *testing.T) {
	s, o := indexedTestStore(t)
	defer s.Close()
	option := search_option.New("chunking strategy", 10)
	outcome := s.SearchWithFallback(option, o)
	assert.True(t, outcome.Degraded)
	assert.Greater(t, 0, float64(len(outcome.Results)))
}
