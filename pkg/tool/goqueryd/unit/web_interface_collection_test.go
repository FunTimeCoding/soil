package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/document_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/collection_page"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web/document_page"
	"net/http"
	"strings"
	"testing"
)

func TestSecondPageOfCollectionLoads(t *testing.T) {
	s := service_tester.New(t).Served()
	c := client_tester.GeneratedClient(t, s)

	for i := range 21 {
		document_tester.Push(
			t,
			c,
			"pushed",
			fmt.Sprintf("note-%02d.md", i),
			fmt.Sprintf("# Note %02d\n\nbody", i),
		)
	}

	first := s.Body(t, collection_page.Path("pushed", "", 1))
	assert.StringContains(t, `href="/collections/pushed?page=2"`, first)
	assert.Integer(
		t,
		http.StatusOK,
		s.Status(t, collection_page.Path("pushed", "", 2)),
	)
	second := s.Body(t, collection_page.Path("pushed", "", 2))
	assert.StringContains(t, `href="/collections/pushed"`, second)
	assert.True(t, !strings.Contains(second, "Not Found"))
	assert.Integer(t, 1, strings.Count(second, `href="/documents/`))
}

func TestMissingDocumentIsNotFoundWithLinkedSuggestions(t *testing.T) {
	s := service_tester.New(t).Served()
	c := client_tester.GeneratedClient(t, s)
	document_tester.Push(t, c, "pushed", "alfa.md", "# Alfa\n\nbody")
	assert.Integer(
		t,
		http.StatusNotFound,
		s.Status(t, document_page.Path("pushed/alfa")),
	)
	page := s.Body(t, document_page.Path("pushed/alfa"))
	assert.StringContains(t, `<a href="/documents/pushed/alfa.md">`, page)
}
