package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/unit/atlassian_tester"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestUnavailableConfluenceIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(
			`{"message":"Confluence is under maintenance"}`,
		),
	)
	var out map[string]any
	e := atlassian_tester.Confluence(t, s.URL).GetPath(
		constant.ConfluenceUser,
		&out,
	)
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "under maintenance", e.Error())
}

func TestRejectedConfluenceTokenStaysFault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
		upstream_server.WithAnswer(
			`{"message":"Basic authentication with passwords is deprecated"}`,
		),
	)
	var out map[string]any
	e := atlassian_tester.Confluence(t, s.URL).GetPath(
		constant.ConfluenceUser,
		&out,
	)
	assert.True(t, unexpected.Is(e))
}
