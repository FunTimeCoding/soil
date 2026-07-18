package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func collectPackageQualifiers(
	all []*packages.Package,
	set *token.FileSet,
	packagePath string,
	oldName string,
	newName string,
) ([]resolve.Reference, string) {
	var result []resolve.Reference

	for _, loaded := range all {
		for ident, o := range loaded.TypesInfo.Uses {
			packageName, okay := o.(*types.PkgName)

			if !okay || packageName.Imported().Path() != packagePath {
				continue
			}

			if packageName.Name() != oldName {
				continue
			}

			scope := loaded.Types.Scope().Innermost(ident.Pos())

			if scope != nil {
				_, shadow := scope.LookupParent(newName, ident.Pos())

				if shadow != nil &&
					shadow.Parent() != types.Universe {
					position := set.Position(ident.Pos())

					return nil, fmt.Sprintf(
						"%s is taken at %s:%d",
						newName,
						position.Filename,
						position.Line,
					)
				}
			}

			result = append(
				result,
				resolve.Reference{Ident: ident, Package: loaded},
			)
		}
	}

	return result, ""
}
