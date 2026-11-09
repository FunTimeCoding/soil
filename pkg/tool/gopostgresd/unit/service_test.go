package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/service_tester"
	"strings"
	"testing"
)

func TestQueryAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Query.Query(t.Context(), "alfa", "SELECT 1")
	assert.FatalOnError(t, e)
	_, f := s.Query.Query(t.Context(), "bravo", "SELECT 1")
	assert.True(t, validation.Is(f))
}

func TestExplainAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Query.Explain(t.Context(), "alfa", "SELECT 1", false)
	assert.FatalOnError(t, e)
	_, f := s.Query.Explain(t.Context(), "bravo", "SELECT 1", false)
	assert.True(t, validation.Is(f))
}

func TestListSchemasAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Catalog.ListSchemas(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	_, f := s.Catalog.ListSchemas(t.Context(), "bravo")
	assert.True(t, unexpected.Is(f))
}

func TestListTablesAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Catalog.ListTables(t.Context(), "alfa", "")
	assert.FatalOnError(t, e)
	_, f := s.Catalog.ListTables(t.Context(), "bravo", "")
	assert.True(t, unexpected.Is(f))
}

func TestDescribeTableAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Catalog.DescribeTable(t.Context(), "alfa", "", "orders")
	assert.FatalOnError(t, e)
	_, f := s.Catalog.DescribeTable(t.Context(), "bravo", "", "orders")
	assert.True(t, unexpected.Is(f))
}

func TestListIndexesAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Catalog.ListIndexes(t.Context(), "alfa", "", "orders")
	assert.FatalOnError(t, e)
	_, f := s.Catalog.ListIndexes(t.Context(), "bravo", "", "orders")
	assert.True(t, unexpected.Is(f))
}

func TestTableSizesAnswersFromNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Catalog.TableSizes(t.Context(), "alfa", "")
	assert.FatalOnError(t, e)
	_, f := s.Catalog.TableSizes(t.Context(), "bravo", "")
	assert.True(t, unexpected.Is(f))
}

func TestCatalogueNamesTravelAsArgumentsNotText(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	c := clients["alfa"]
	_, e := s.Catalog.DescribeTable(t.Context(), "alfa", "sales", "it's")
	assert.FatalOnError(t, e)
	v := c.Statements()[0]
	assert.False(t, strings.Contains(v.SQL, "sales"))
	assert.False(t, strings.Contains(v.SQL, "it's"))
	assert.Count(t, 2, v.Arguments)
	assert.Any(t, "sales", v.Arguments[0])
	assert.Any(t, "it's", v.Arguments[1])
}

func TestEmptySchemaMeansPublic(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	c := clients["alfa"]
	_, e := s.Catalog.ListTables(t.Context(), "alfa", "")
	assert.FatalOnError(t, e)
	assert.Any(t, "public", c.Statements()[0].Arguments[0])
}

func TestExplainPrefixesStatement(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	c := clients["alfa"]
	_, e := s.Query.Explain(t.Context(), "alfa", "SELECT 1", false)
	assert.FatalOnError(t, e)
	_, f := s.Query.Explain(t.Context(), "alfa", "SELECT 1", true)
	assert.FatalOnError(t, f)
	assert.String(t, "EXPLAIN SELECT 1", c.Statements()[0].SQL)
	assert.String(t, "EXPLAIN ANALYZE SELECT 1", c.Statements()[1].SQL)
}

func TestUnreachableInstanceFailsFastWithoutBeingAskedAgain(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	c := clients["alfa"]
	c.Failure = refusedConnection()
	_, e := s.Query.Query(t.Context(), "alfa", "SELECT 1")
	assert.True(t, unavailable.Is(e))
	c.Failure = nil
	_, f := s.Query.Query(t.Context(), "alfa", "SELECT 1")
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "postgres:alfa unreachable since", f.Error())
	assert.Count(t, 1, c.Statements())
}

func TestOneUnreachableInstanceLeavesOthersAnswering(t *testing.T) {
	s, clients := service_tester.New(t, "alfa", "bravo")
	clients["alfa"].Failure = refusedConnection()
	_, e := s.Query.Query(t.Context(), "alfa", "SELECT 1")
	assert.True(t, unavailable.Is(e))
	_, f := s.Query.Query(t.Context(), "bravo", "SELECT 1")
	assert.FatalOnError(t, f)
}
