//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestClearBindings(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	s.BindModelContextSession(r.Callsign, "mcp-session-abc")
	e := s.GetSession("session-1")
	assert.String(t, "mcp-session-abc", e.ModelContextSession)
	s.Store.ClearBindings()
	e = s.GetSession("session-1")
	assert.String(t, "", e.ModelContextSession)
}
