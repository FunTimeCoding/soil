//go:build local

package cross_service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/cross_service_tester"
	goquerydConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

func TestProfileShowsCompletions(t *testing.T) {
	s := cross_service_tester.New(t)
	s.Goqueryd.MustCallTool(
		goquerydConstant.Push,
		map[string]any{
			goquerydConstant.Collection: "completions",
			goquerydConstant.Path:       "test-session/1",
			goquerydConstant.Body:       "built the search pipeline",
			goquerydConstant.Metadata: map[string]string{
				"source_type":  "session-completion",
				"session_name": "test-session",
			},
		},
	)
	result := s.Gomemoryd.MustCallTool(
		constant.Profile,
		map[string]any{},
	)
	assert.StringContains(t, "test-session", result)
	assert.StringContains(t, "built the search pipeline", result)
}

func TestProfileShowsMultipleCompletionsPerSession(t *testing.T) {
	s := cross_service_tester.New(t)
	s.Goqueryd.MustCallTool(
		goquerydConstant.Push,
		map[string]any{
			goquerydConstant.Collection: "completions",
			goquerydConstant.Path:       "test-session/1",
			goquerydConstant.Body:       "built the API",
			goquerydConstant.Metadata: map[string]string{
				"source_type":  "session-completion",
				"session_name": "test-session",
			},
		},
	)
	s.Goqueryd.MustCallTool(
		goquerydConstant.Push,
		map[string]any{
			goquerydConstant.Collection: "completions",
			goquerydConstant.Path:       "test-session/2",
			goquerydConstant.Body:       "wrote the tests",
			goquerydConstant.Metadata: map[string]string{
				"source_type":  "session-completion",
				"session_name": "test-session",
			},
		},
	)
	result := s.Gomemoryd.MustCallTool(
		constant.Profile,
		map[string]any{},
	)
	assert.StringContains(t, "built the API", result)
	assert.StringContains(t, "wrote the tests", result)
}

func TestProfileCompletionsEmptyWhenNone(t *testing.T) {
	s := cross_service_tester.New(t)
	result := s.Gomemoryd.MustCallTool(
		constant.Profile,
		map[string]any{},
	)
	assert.StringNotContains(t, "completions", result)
}
