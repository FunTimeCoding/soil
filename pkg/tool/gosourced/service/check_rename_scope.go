package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkRenameScope(
	all []*packages.Package,
	set *token.FileSet,
	packagePath string,
	newName string,
	references []resolve_reference.Reference,
) string {
	if types.Universe.Lookup(newName) != nil {
		return fmt.Sprintf("%s is a built-in name", newName)
	}

	for _, loaded := range all {
		if loaded.PkgPath != packagePath {
			continue
		}

		for _, file := range loaded.Syntax {
			if o := loaded.TypesInfo.Scopes[file].Lookup(newName); o != nil {
				return takenAt(set, newName, o.Pos())
			}
		}
	}

	for _, f := range references {
		if f.Package.PkgPath != packagePath {
			continue
		}

		scope := f.Package.Types.Scope().Innermost(f.Ident.Pos())

		if scope == nil {
			continue
		}

		_, shadow := scope.LookupParent(newName, f.Ident.Pos())

		if shadow != nil && shadow.Parent() != types.Universe {
			return takenAt(set, newName, f.Ident.Pos())
		}
	}

	return ""
}
