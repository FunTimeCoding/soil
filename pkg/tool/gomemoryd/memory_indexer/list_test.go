package memory_indexer_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/memory_indexer_tester"
	"testing"
)

func TestListReturnsResults(t *testing.T) {
	s := memory_indexer_tester.New(
		t,
		memory_indexer_tester.EmptyFixture(),
		memory_indexer_tester.ListFixture(
			memory_indexer_tester.TestResult(
				"test-session/1",
				"built the API",
				0,
			),
			memory_indexer_tester.TestResult(
				"test-session/2",
				"wrote the tests",
				0,
			),
		),
	)
	outcome, e := s.List(
		"completions",
		map[string]string{"source_type": "session-completion"},
		5,
		0,
		true,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, outcome.Results)
	assert.String(t, "test-session/1", outcome.Results[0].Path)
	assert.String(t, "built the API", outcome.Results[0].Body)
}

func TestListEmptyCollection(t *testing.T) {
	s := memory_indexer_tester.New(
		t,
		memory_indexer_tester.EmptyFixture(),
		memory_indexer_tester.ListFixture(),
	)
	outcome, e := s.List("completions", nil, 10, 0, false)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, outcome.Results)
}

func TestListWithoutBody(t *testing.T) {
	s := memory_indexer_tester.New(
		t,
		memory_indexer_tester.EmptyFixture(),
		memory_indexer_tester.ListFixture(
			memory_indexer_tester.TestResultNoBody("test-session/1"),
			memory_indexer_tester.TestResultNoBody("test-session/2"),
		),
	)
	outcome, e := s.List("completions", nil, 5, 0, false)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, outcome.Results)
	assert.String(t, "", outcome.Results[0].Body)
}
