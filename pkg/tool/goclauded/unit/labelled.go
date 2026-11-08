package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func labelled(
	t *testing.T,
	s *service_tester.Tester,
	key string,
) string {
	t.Helper()
	value, e := s.Label.GetLabel("session-1", key)
	assert.FatalOnError(t, e)

	return value
}
