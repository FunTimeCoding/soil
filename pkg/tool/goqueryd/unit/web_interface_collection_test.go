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

func TestSecondPageOfCollectionLoads(t *testing.T) {
	s, c := servedWeb(t)

	for i := range 21 {
		pushDocument(
			t,
			c,
			"pushed",
			fmt.Sprintf("note-%02d.md", i),
			fmt.Sprintf("# Note %02d\n\nbody", i),
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

func TestMissingDocumentIsNotFoundWithLinkedSuggestions(t *testing.T) {
	s, c := servedWeb(t)
	pushDocument(t, c, "pushed", "alfa.md", "# Alfa\n\nbody")
	status, page := servedPage(t, s, document_page.Path("pushed/alfa"))
	assert.Integer(t, http.StatusNotFound, status)
	assert.StringContains(t, `<a href="/documents/pushed/alfa.md">`, page)
}
