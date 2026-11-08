package census

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"

func methodResult(
	methods map[string]*reading,
	name string,
	stores map[string]*storeReading,
	tables map[string]string,
) *result.CensusMethod {
	r := methods[name]
	m := result.NewCensusMethod(r.name, r.file)
	m.Callers = callerResults(r.callers)
	m.Calls = sortedKeys(r.calls)
	m.Functions = sortedKeys(r.functions)
	m.Fields = sortedKeys(r.fields)
	m.Uses = useResults(r.uses)
	m.Packages = sortedKeys(r.packages)
	m.Atomically = r.atomically
	m.TransactionStep = r.step
	m.Tables = tableResult(stores, tables, []*reading{r})
	m.Reach = tableResult(stores, tables, reached(methods, name))

	return m
}
