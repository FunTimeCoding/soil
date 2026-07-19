package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/memory_indexer_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face/search_option"
	"testing"
)

func TestSearchReturnsResults(t *testing.T) {
	s := memory_indexer_tester.New(
		t,
		memory_indexer_tester.SearchFixture(
			memory_indexer_tester.TestResult(
				"memory/1",
				"captureFail content",
				0.93,
			),
			memory_indexer_tester.TestResult(
				"memory/2",
				"error handling",
				0.71,
			),
		),
		memory_indexer_tester.EmptyFixture(),
	)
	results, e := s.Search(
		search_option.New("error handling", "memories", 10),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, results)
	assert.String(t, "memory/1", results[0].Path)
	assert.StringContains(t, "captureFail", results[0].Body)
}

func TestSearchEmptyResults(t *testing.T) {
	s := memory_indexer_tester.New(
		t,
		memory_indexer_tester.SearchFixture(),
		memory_indexer_tester.EmptyFixture(),
	)
	results, e := s.Search(
		search_option.New("nonexistent", "memories", 10),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, results)
}
