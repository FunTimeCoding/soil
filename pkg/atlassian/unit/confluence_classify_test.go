package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestAnUnavailableConfluenceIsUnavailable(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusServiceUnavailable,
		`{"message":"Confluence is under maintenance"}`,
	)
	var out map[string]any
	e := newConfluenceClient(t, s.URL).GetPath(constant.ConfluenceUser, &out)
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "under maintenance", e.Error())
}

func TestARejectedConfluenceTokenStaysAFault(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusUnauthorized,
		`{"message":"Basic authentication with passwords is deprecated"}`,
	)
	var out map[string]any
	e := newConfluenceClient(t, s.URL).GetPath(constant.ConfluenceUser, &out)
	assert.True(t, unexpected.Is(e))
}
