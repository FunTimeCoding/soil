package coordination

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"net/http"
	"testing"
)

func TestShortIdentifierSurvivesDisplayPaths(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	s.Store.EnsureSession("tiny")
	s.Store.Store.UpdateFields(
		"tiny",
		map[string]any{"turn_count": 4, "lines": 20},
	)
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "short identifier")
	a.MustCallTool(constant.ListSessions, map[string]any{})
	assert.Integer(t, http.StatusOK, pageStatus(t, s, "/sessions/tiny"))
	assert.Integer(t, http.StatusOK, pageStatus(t, s, constant.SessionsPath))
}

func pageStatus(
	t *testing.T,
	s *base.Server,
	path string,
) int {
	t.Helper()
	response, e := http.Get(fmt.Sprintf("http://localhost:%d%s", s.Port, path))
	assert.FatalOnError(t, e)

	defer errors.PanicClose(response.Body)

	return response.StatusCode
}
