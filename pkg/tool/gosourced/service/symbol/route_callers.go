package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

func routeCallers(
	all []*packages.Package,
	set *token.FileSet,
	p *packages.Package,
	source *types.Named,
	destination *types.Named,
	entries []*relocation.Entry,
) (map[*ast.SelectorExpr]string, string) {
	result := map[*ast.SelectorExpr]string{}
	var field *types.Var

	for _, entry := range entries {
		pointer := isPointerReceiver(entry.Object)

		for _, f := range resolve.FindAllReferences(all, entry.Object) {
			if movedContains(entries, f.Ident.Pos()) {
				continue
			}

			position := set.Position(f.Ident.Pos())
			at := fmt.Sprintf(
				"%s is called at %s:%d",
				entry.Symbol,
				position.Filename,
				position.Line,
			)

			if field == nil {
				field = holderField(source, destination)
			}

			if field == nil {
				return nil, fmt.Sprintf(
					"%s and %s holds no single field of type %s to call it through",
					at,
					qualifiedType(source),
					qualifiedType(destination),
				)
			}

			file := lookup.SyntaxFile(f.Package, f.Ident.Pos())
			path, _ := astutil.PathEnclosingInterval(
				file,
				f.Ident.Pos(),
				f.Ident.End(),
			)
			var s *ast.SelectorExpr

			if len(path) > 1 {
				s, _ = path[1].(*ast.SelectorExpr)
			}

			if s == nil || s.Sel != f.Ident {
				return nil, join.Empty(
					at,
					" not through a selector - move it by hand",
				)
			}

			selection := f.Package.TypesInfo.Selections[s]

			if selection == nil || selection.Kind() != types.MethodVal {
				return nil, join.Empty(
					at,
					" as a method expression - move it by hand",
				)
			}

			if len(selection.Index()) != 1 {
				return nil, join.Empty(
					at,
					" through an embedding - move it by hand",
				)
			}

			if !field.Exported() && f.Package.PkgPath != p.PkgPath {
				return nil, fmt.Sprintf(
					"%s, where the unexported %s cannot be reached",
					at,
					field.Name(),
				)
			}

			_, recvPointer := selection.Recv().(*types.Pointer)
			_, fieldPointer := field.Type().(*types.Pointer)

			if pointer && !fieldPointer && !recvPointer {
				return nil, fmt.Sprintf(
					"%s on a value, where %s.%s may not be addressable - move it by hand",
					at,
					qualifiedType(source),
					field.Name(),
				)
			}

			result[s] = field.Name()
		}
	}

	return result, ""
}
