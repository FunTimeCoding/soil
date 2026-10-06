package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	docker "github.com/funtimecoding/soil/pkg/docker/constant"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"testing"
)

func TestTagsReadsThePage(t *testing.T) {
	s, last := upstream_tester.New(
		t,
		http.StatusOK,
		`{"results":[{"name":"1.27-alpine"}]}`,
	)
	tags, e := newHubClient(t, s.URL).Tags("library/alfa")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, tags)
	assert.String(t, "1.27-alpine", tags[0].Name)
	r := last()
	assert.String(t, "/v2/repositories/library/alfa/tags", r.URL.Path)
	assert.String(t, "100", r.URL.Query().Get(docker.PageSizeParameter))
	assert.String(t, "application/json", r.Header.Get(constant.Accept))
}

func TestMissingImageIsNotFound(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusNotFound,
		`{"message":"object not found"}`,
	)
	_, e := newHubClient(t, s.URL).Tags("library/bravo")
	assert.True(t, not_found.Is(e))
}
