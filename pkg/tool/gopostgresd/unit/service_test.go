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

func TestQueryAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Query(t.Context(), "alfa", "SELECT 1")
	assert.FatalOnError(t, e)
	_, f := s.Query(t.Context(), "bravo", "SELECT 1")
	assert.True(t, validation.Is(f))
}

func TestExplainAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.Explain(t.Context(), "alfa", "SELECT 1", false)
	assert.FatalOnError(t, e)
	_, f := s.Explain(t.Context(), "bravo", "SELECT 1", false)
	assert.True(t, validation.Is(f))
}

func TestListSchemasAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.ListSchemas(t.Context(), "alfa")
	assert.FatalOnError(t, e)
	_, f := s.ListSchemas(t.Context(), "bravo")
	assert.True(t, unexpected.Is(f))
}

func TestListTablesAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.ListTables(t.Context(), "alfa", "")
	assert.FatalOnError(t, e)
	_, f := s.ListTables(t.Context(), "bravo", "")
	assert.True(t, unexpected.Is(f))
}

func TestDescribeTableAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.DescribeTable(t.Context(), "alfa", "", "orders")
	assert.FatalOnError(t, e)
	_, f := s.DescribeTable(t.Context(), "bravo", "", "orders")
	assert.True(t, unexpected.Is(f))
}

func TestListIndexesAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.ListIndexes(t.Context(), "alfa", "", "orders")
	assert.FatalOnError(t, e)
	_, f := s.ListIndexes(t.Context(), "bravo", "", "orders")
	assert.True(t, unexpected.Is(f))
}

func TestTableSizesAnswersFromTheNamedInstance(t *testing.T) {
	s := refusingBravo(t)
	_, e := s.TableSizes(t.Context(), "alfa", "")
	assert.FatalOnError(t, e)
	_, f := s.TableSizes(t.Context(), "bravo", "")
	assert.True(t, unexpected.Is(f))
}

func TestCatalogueNamesTravelAsArgumentsNotText(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	c := clients["alfa"]
	_, e := s.DescribeTable(t.Context(), "alfa", "sales", "it's")
	assert.FatalOnError(t, e)
	v := c.Statements()[0]
	assert.False(t, strings.Contains(v.SQL, "sales"))
	assert.False(t, strings.Contains(v.SQL, "it's"))
	assert.Count(t, 2, v.Arguments)
	assert.Any(t, "sales", v.Arguments[0])
	assert.Any(t, "it's", v.Arguments[1])
}

func TestAnEmptySchemaMeansPublic(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	c := clients["alfa"]
	_, e := s.ListTables(t.Context(), "alfa", "")
	assert.FatalOnError(t, e)
	assert.Any(t, "public", c.Statements()[0].Arguments[0])
}

func TestExplainPrefixesTheStatement(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	c := clients["alfa"]
	_, e := s.Explain(t.Context(), "alfa", "SELECT 1", false)
	assert.FatalOnError(t, e)
	_, f := s.Explain(t.Context(), "alfa", "SELECT 1", true)
	assert.FatalOnError(t, f)
	assert.String(t, "EXPLAIN SELECT 1", c.Statements()[0].SQL)
	assert.String(t, "EXPLAIN ANALYZE SELECT 1", c.Statements()[1].SQL)
}

func TestAnUnreachableInstanceFailsFastWithoutBeingAskedAgain(t *testing.T) {
	s, clients := service_tester.New(t, "alfa")
	c := clients["alfa"]
	c.Failure = refusedConnection()
	_, e := s.Query(t.Context(), "alfa", "SELECT 1")
	assert.True(t, unavailable.Is(e))
	c.Failure = nil
	_, f := s.Query(t.Context(), "alfa", "SELECT 1")
	assert.True(t, unavailable.Is(f))
	assert.StringContains(t, "postgres:alfa unreachable since", f.Error())
	assert.Count(t, 1, c.Statements())
}

func TestOneUnreachableInstanceLeavesTheOthersAnswering(t *testing.T) {
	s, clients := service_tester.New(t, "alfa", "bravo")
	clients["alfa"].Failure = refusedConnection()
	_, e := s.Query(t.Context(), "alfa", "SELECT 1")
	assert.True(t, unavailable.Is(e))
	_, f := s.Query(t.Context(), "bravo", "SELECT 1")
	assert.FatalOnError(t, f)
}
