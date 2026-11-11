package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"go/ast"
	"go/types"
	"path/filepath"
	"strings"
)

func storeReadings(
	directory string,
	module string,
	keys map[string]bool,
	missing unloaded,
) (map[string]*storeReading, error) {
	result := map[string]*storeReading{}
	loaded := map[string]bool{}
	pending := map[string]bool{}

	for k := range keys {
		pending[keyPackage(k)] = true
	}

	for len(pending) > 0 {
		paths := sortedKeys(pending)
		pending = map[string]bool{}

		for _, path := range paths {
			loaded[path] = true
		}

		all, broken, _, e := load.Lenient(directory, paths...)

		if e != nil {
			return nil, e
		}

		noteUnloaded(missing, broken, constant.CensusLostTables)

		for _, p := range all {
			for _, f := range p.Syntax {
				file := filepath.Base(p.Fset.Position(f.Pos()).Filename)

				if strings.HasSuffix(file, constant.TestFileSuffix) {
					continue
				}

				for _, d := range f.Decls {
					x, okay := d.(*ast.FuncDecl)

					if !okay || x.Body == nil || x.Recv == nil {
						continue
					}

					o, okay := p.TypesInfo.Defs[x.Name].(*types.Func)

					if !okay || owner(o) == nil {
						continue
					}

					s := readStore(p, x, module)
					result[methodKey(owner(o), x.Name.Name)] = s

					for c := range s.calls {
						if path := keyPackage(c); !loaded[path] {
							pending[path] = true
						}
					}
				}
			}
		}
	}

	return result, nil
}
