package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/census/reading"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
)

func tableResult(
	stores map[string]*reading.Store,
	tables map[string]string,
	readings []*reading.Reading,
) *result.CensusTables {
	read := map[string]bool{}
	written := map[string]bool{}

	for _, r := range readings {
		for _, k := range sortedKeys(r.Keys) {
			tableClosure(stores, tables, k, read, written)
		}
	}

	return result.NewCensusTables(sortedKeys(read), sortedKeys(written))
}
