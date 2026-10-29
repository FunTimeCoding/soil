package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"testing"
)

func TestTheClientRoundTripsADocument(t *testing.T) {
	c := servedClient(t)
	x := t.Context()
	metadata := map[string][]string{"scope": {"alfa"}}
	_, e := c.PostDocument(
		x,
		client.PostDocumentJSONRequestBody{
			Collection: "notes",
			Path:       "alfa.md",
			Body:       "# Alfa\n\nzulu quasar.\n",
			Metadata:   &metadata,
		},
	)
	assert.FatalOnError(t, e)
	document, f := c.GetDocument(
		x,
		&client.GetDocumentParams{Path: "notes/alfa.md"},
	)
	assert.FatalOnError(t, f)
	assert.String(t, "Alfa", document.Title)
	limit := 10
	mode := client.Keyword
	found, g := c.GetSearch(
		x,
		&client.GetSearchParams{
			Query:      "quasar",
			Collection: new("notes"),
			Limit:      &limit,
			Mode:       &mode,
		},
	)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, found.Results)
	listed, h := c.GetList(
		x,
		&client.GetListParams{
			Collection: "notes",
			Metadata:   &map[string]string{"scope": "alfa"},
		},
	)
	assert.FatalOnError(t, h)
	assert.Count(t, 1, listed.Results)
	assert.FatalOnError(
		t,
		c.PostContext(
			x,
			client.PostContextJSONRequestBody{
				Collection:  "notes",
				PathPrefix:  "/",
				Description: "notes",
			},
		),
	)
	contexts, i := c.GetContext(x)
	assert.FatalOnError(t, i)
	assert.Count(t, 1, contexts)
	deleted, j := c.DeleteDocument(
		x,
		&client.DeleteDocumentParams{Collection: "notes", Path: "alfa.md"},
	)
	assert.FatalOnError(t, j)
	assert.True(t, deleted)
	again, k := c.DeleteDocument(
		x,
		&client.DeleteDocumentParams{Collection: "notes", Path: "alfa.md"},
	)
	assert.FatalOnError(t, k)
	assert.False(t, again)
	status, l := c.GetStatus(x)
	assert.FatalOnError(t, l)
	assert.Integer(t, 0, status.TotalDocuments)
	assert.FatalOnError(t, c.Reachable(x))
}

func TestTheClientAnswersEachClass(t *testing.T) {
	c := servedClient(t)
	x := t.Context()
	_, e := c.GetDocument(
		x,
		&client.GetDocumentParams{Path: "notes/missing.md"},
	)
	assert.True(t, not_found.Is(e))
	assert.StringContains(t, "document not found", e.Error())
	_, f := c.GetDocument(x, &client.GetDocumentParams{Path: "noslash"})
	assert.True(t, validation.Is(f))
}
