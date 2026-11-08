package census

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"

func tableResult(
	stores map[string]*storeReading,
	tables map[string]string,
	readings []*reading,
) *result.CensusTables {
	read := map[string]bool{}
	written := map[string]bool{}

	for _, r := range readings {
		for _, k := range sortedKeys(r.keys) {
			tableClosure(stores, tables, k, read, written)
		}
	}

	return result.NewCensusTables(sortedKeys(read), sortedKeys(written))
}
