package service

import (
	"go/token"
	"golang.org/x/tools/go/packages"
	"sort"
	"strings"
)

func moduleSymbols(
	all []*packages.Package,
	set *token.FileSet,
	modulePath string,
) []*ModuleSymbol {
	seen := make(map[string]*ModuleSymbol)

	for _, loaded := range all {
		if strings.HasSuffix(loaded.ID, ".test") ||
			loaded.TypesInfo == nil ||
			belongsToModule(loaded.PkgPath, modulePath) {
			continue
		}

		for expression, selection := range loaded.TypesInfo.Selections {
			o := selection.Obj()

			if !objectBelongsToModule(o, modulePath) {
				continue
			}

			addModuleSymbol(
				seen,
				&ModuleSymbol{
					PackagePath: o.Pkg().Path(),
					Owner:       selectionOwner(selection),
					Name:        o.Name(),
					Object:      o,
					Position:    set.Position(expression.Pos()),
				},
			)
		}

		for identity, o := range loaded.TypesInfo.Uses {
			if !objectBelongsToModule(o, modulePath) ||
				o.Parent() != o.Pkg().Scope() {
				continue
			}

			addModuleSymbol(
				seen,
				&ModuleSymbol{
					PackagePath: o.Pkg().Path(),
					Name:        o.Name(),
					Object:      o,
					Position:    set.Position(identity.Pos()),
				},
			)
		}
	}

	var keys []string

	for key := range seen {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	var result []*ModuleSymbol

	for _, key := range keys {
		result = append(result, seen[key])
	}

	return result
}
