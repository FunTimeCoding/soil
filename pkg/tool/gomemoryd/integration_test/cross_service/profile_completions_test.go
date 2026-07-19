//go:build local

package cross_service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/cross_service_tester"
	query "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

func TestProfileShowsCompletions(t *testing.T) {
	s := cross_service_tester.New(t)
	s.QueryClient.MustCallTool(
		query.Push,
		map[string]any{
			query.Collection: "completions",
			query.Path:       "test-session/1",
			query.Body:       "built the search pipeline",
			query.Metadata: map[string]string{
				"source_type":  "session-completion",
				"session_name": "test-session",
			},
		},
	)
	result := s.MemoryClient.MustCallTool(
		constant.Profile,
		map[string]any{},
	)
	assert.StringContains(t, "test-session", result)
	assert.StringContains(t, "built the search pipeline", result)
}

func TestProfileShowsMultipleCompletionsPerSession(t *testing.T) {
	s := cross_service_tester.New(t)
	s.QueryClient.MustCallTool(
		query.Push,
		map[string]any{
			query.Collection: "completions",
			query.Path:       "test-session/1",
			query.Body:       "built the API",
			query.Metadata: map[string]string{
				"source_type":  "session-completion",
				"session_name": "test-session",
			},
		},
	)
	s.QueryClient.MustCallTool(
		query.Push,
		map[string]any{
			query.Collection: "completions",
			query.Path:       "test-session/2",
			query.Body:       "wrote the tests",
			query.Metadata: map[string]string{
				"source_type":  "session-completion",
				"session_name": "test-session",
			},
		},
	)
	result := s.MemoryClient.MustCallTool(
		constant.Profile,
		map[string]any{},
	)
	assert.StringContains(t, "built the API", result)
	assert.StringContains(t, "wrote the tests", result)
}

func TestProfileCompletionsEmptyWhenNone(t *testing.T) {
	assert.StringNotContains(
		t,
		"completions",
		cross_service_tester.New(t).MemoryClient.MustCallTool(
			constant.Profile,
			map[string]any{},
		),
	)
}
