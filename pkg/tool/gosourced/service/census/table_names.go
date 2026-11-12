package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/census/reading"
	"go/types"
	"gorm.io/gorm/schema"
)

func tableNames(
	directory string,
	stores map[string]*reading.Store,
	missing unloaded,
) (map[string]string, error) {
	result := map[string]string{}
	declared := map[string]*types.Named{}
	paths := map[string]bool{}
	naming := schema.NamingStrategy{}

	for _, s := range stores {
		for q, m := range s.Models {
			if _, done := result[q]; done || declared[q] != nil {
				continue
			}

			if types.NewMethodSet(m).Lookup(m.Obj().Pkg(), constant.TableNameMethod) == nil &&
				types.NewMethodSet(types.NewPointer(m)).Lookup(
					m.Obj().Pkg(),
					constant.TableNameMethod,
				) == nil {
				result[q] = naming.TableName(m.Obj().Name())

				continue
			}

			declared[q] = m
			paths[m.Obj().Pkg().Path()] = true
		}
	}

	if len(declared) == 0 {
		return result, nil
	}

	all, broken, _, e := load.Lenient(directory, sortedKeys(paths)...)

	if e != nil {
		return nil, e
	}

	noteUnloaded(missing, broken, constant.CensusLostTables)

	for q, m := range declared {
		result[q] = declaredTable(all, m)
	}

	return result, nil
}
