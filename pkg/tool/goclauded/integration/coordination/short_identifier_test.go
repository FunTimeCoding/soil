package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"net/http"
	"testing"
)

func TestShortIdentifierSurvivesDisplayPaths(t *testing.T) {
	s := base.New(t)
	s.Store.EnsureSession("tiny")
	s.Store.Store.UpdateFields(
		"tiny",
		map[string]any{"turn_count": 4, "lines": 20},
	)
	a := s.NewSession(t)
	a.Announce(a.Name(), "short identifier")
	a.MustCallTool(constant.ListSessions, map[string]any{})
	assert.Integer(t, http.StatusOK, pageStatus(t, s, "/sessions/tiny"))
	assert.Integer(t, http.StatusOK, pageStatus(t, s, constant.SessionsPath))
}
