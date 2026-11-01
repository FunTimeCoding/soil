package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/collection_page"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/document_page"
	"net/http"
	"strings"
	"testing"
)

func TestTheSecondPageOfACollectionLoads(t *testing.T) {
	s, v := servedWeb(t)

	for i := range 21 {
		v.PushDocument(
			t.Context(),
			"pushed",
			fmt.Sprintf("note-%02d.md", i),
			fmt.Sprintf("# Note %02d\n\nbody", i),
			nil,
		)
	}

	_, first := servedPage(t, s, collection_page.Path("pushed", "", 1))
	assert.StringContains(t, `href="/collections/pushed?page=2"`, first)
	status, second := servedPage(t, s, collection_page.Path("pushed", "", 2))
	assert.Integer(t, http.StatusOK, status)
	assert.StringContains(t, `href="/collections/pushed"`, second)
	assert.True(t, !strings.Contains(second, "Not Found"))
	assert.Integer(t, 1, strings.Count(second, `href="/documents/`))
}

func TestAMissingDocumentIsNotFoundWithLinkedSuggestions(t *testing.T) {
	s, v := servedWeb(t)
	v.PushDocument(t.Context(), "pushed", "alfa.md", "# Alfa\n\nbody", nil)
	status, page := servedPage(t, s, document_page.Path("pushed/alfa"))
	assert.Integer(t, http.StatusNotFound, status)
	assert.StringContains(t, `<a href="/documents/pushed/alfa.md">`, page)
}
