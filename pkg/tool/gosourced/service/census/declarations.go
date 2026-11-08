package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
)

func declarations(
	p *packages.Package,
	subject *types.Named,
	module string,
) (map[string]*reading, map[string]*reading) {
	stores := fieldTypes(subject)
	methods := map[string]*reading{}
	functions := map[string]*reading{}

	for _, f := range p.Syntax {
		file := filepath.Base(p.Fset.Position(f.Pos()).Filename)

		if strings.HasSuffix(file, constant.TestFileSuffix) {
			continue
		}

		for _, d := range f.Decls {
			x, okay := d.(*ast.FuncDecl)

			if !okay || x.Body == nil {
				continue
			}

			if x.Recv == nil {
				functions[x.Name.Name] = readBody(
					p,
					x,
					file,
					subject,
					module,
					stores,
				)

				continue
			}

			o, okay := p.TypesInfo.Defs[x.Name].(*types.Func)

			if okay && owner(o) != nil && sameType(owner(o), subject) {
				methods[x.Name.Name] = readBody(
					p,
					x,
					file,
					subject,
					module,
					stores,
				)
			}
		}
	}

	return methods, functions
}
