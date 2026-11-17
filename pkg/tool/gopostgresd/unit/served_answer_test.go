package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/service_tester"
	"net/http"
	"testing"
)

func TestRestTypoInTheirQueryIsQuiet(t *testing.T) {
	b := base.New(t)
	b.Alfa().Failure = service_tester.SyntaxRefusal()
	c := b.GeneratedClient()
	v, e := c.QueryWithResponse(
		t.Context(),
		client.QueryJSONRequestBody{Instance: new("alfa"), Sql: "SELEC 1"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusBadRequest, v.StatusCode())
	assert.Count(t, 0, b.Reporter.Events())
}

func TestRestRefusalOfOurCatalogueIsCaptured(t *testing.T) {
	b := base.New(t)
	b.Alfa().Failure = service_tester.SyntaxRefusal()
	c := b.GeneratedClient()
	v, e := c.ListSchemasWithResponse(
		t.Context(),
		&client.ListSchemasParams{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, v.StatusCode())
	assert.Count(t, 1, b.Reporter.Events())
}

func TestModelContextTypoInTheirQueryIsQuiet(t *testing.T) {
	b := base.New(t)
	b.Alfa().Failure = service_tester.SyntaxRefusal()
	m := b.Tools()
	assert.StringContains(
		t,
		"syntax error",
		m.MustCallToolError(constant.Query, map[string]any{"sql": "SELEC 1"}),
	)
	assert.Count(t, 0, b.Reporter.Events())
}

func TestModelContextRefusalOfOurCatalogueIsCaptured(t *testing.T) {
	b := base.New(t)
	b.Alfa().Failure = service_tester.SyntaxRefusal()
	m := b.Tools()
	m.MustCallToolError(constant.ListSchemas, nil)
	assert.Count(t, 1, b.Reporter.Events())
}
