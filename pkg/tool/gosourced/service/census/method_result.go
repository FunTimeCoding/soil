package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/census/reading"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
)

func methodResult(
	methods map[string]*reading.Reading,
	name string,
	stores map[string]*reading.Store,
	tables map[string]string,
) *result.CensusMethod {
	r := methods[name]
	m := result.NewCensusMethod(r.Name, r.File)
	m.Callers = callerResults(r.Callers)
	m.Calls = sortedKeys(r.Calls)
	m.Functions = sortedKeys(r.Functions)
	m.Fields = sortedKeys(r.Fields)
	m.Uses = useResults(r.Uses)
	m.Packages = sortedKeys(r.Packages)
	m.Atomically = r.Atomically
	m.TransactionStep = r.Step
	m.Tables = tableResult(stores, tables, []*reading.Reading{r})
	m.Reach = tableResult(stores, tables, reached(methods, name))

	return m
}
