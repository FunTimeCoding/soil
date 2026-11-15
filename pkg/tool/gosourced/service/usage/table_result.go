package usage

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/usage/reading"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
)

func tableResult(
	stores map[string]*reading.Store,
	tables map[string]string,
	readings []*reading.Reading,
) *result.UsageTables {
	read := map[string]bool{}
	written := map[string]bool{}

	for _, r := range readings {
		for _, k := range sortedKeys(r.Keys) {
			tableClosure(stores, tables, k, read, written)
		}
	}

	return result.NewUsageTables(sortedKeys(read), sortedKeys(written))
}
