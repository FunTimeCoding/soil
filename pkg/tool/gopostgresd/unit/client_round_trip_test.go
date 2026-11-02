package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/client"
	generated "github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestTheClientReadsColumnsAndRowsThroughTheMount(t *testing.T) {
	s, _, _, m := served(t, nil)
	m.SeedResult(
		"SELECT a.id, b.id",
		result.New([]string{"id", "id"}, [][]any{{1.0, 7.0}}),
	)
	c := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(c.Close)
	v, e := c.Query(
		t.Context(),
		generated.QueryJSONRequestBody{
			Instance: new("alfa"),
			Sql:      "SELECT a.id, b.id",
		},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v.Columns)
	assert.Any(t, 7.0, v.Rows[0][1])
}

func TestTheClientClassifiesARefusalThroughTheMount(t *testing.T) {
	s, _, _, _ := served(t, syntaxRefusal())
	c := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(c.Close)
	_, e := c.Query(
		t.Context(),
		generated.QueryJSONRequestBody{Instance: new("alfa"), Sql: "SELEC 1"},
	)
	assert.True(t, validation.Is(e))
}
