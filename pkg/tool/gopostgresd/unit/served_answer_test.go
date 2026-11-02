package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	"net/http"
	"testing"
)

func TestARestTypoInTheirQueryIsQuiet(t *testing.T) {
	_, c, r, _ := served(t, syntaxRefusal())
	v, e := c.QueryWithResponse(
		t.Context(),
		client.QueryJSONRequestBody{Instance: new("alfa"), Sql: "SELEC 1"},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusBadRequest, v.StatusCode())
	assert.Count(t, 0, r.Events())
}

func TestARestRefusalOfOurCatalogueIsCaptured(t *testing.T) {
	_, c, r, _ := served(t, syntaxRefusal())
	v, e := c.ListSchemasWithResponse(
		t.Context(),
		&client.ListSchemasParams{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusInternalServerError, v.StatusCode())
	assert.Count(t, 1, r.Events())
}

func TestAModelContextTypoInTheirQueryIsQuiet(t *testing.T) {
	s, _, r, _ := served(t, syntaxRefusal())
	m := s.Connect(t)
	defer m.Close()
	assert.StringContains(
		t,
		"syntax error",
		m.MustCallToolError(constant.Query, map[string]any{"sql": "SELEC 1"}),
	)
	assert.Count(t, 0, r.Events())
}

func TestAModelContextRefusalOfOurCatalogueIsCaptured(t *testing.T) {
	s, _, r, _ := served(t, syntaxRefusal())
	m := s.Connect(t)
	defer m.Close()
	m.MustCallToolError(constant.ListSchemas, nil)
	assert.Count(t, 1, r.Events())
}
