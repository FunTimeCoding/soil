package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module_symbol"
	"go/token"
	"golang.org/x/tools/go/packages"
	"sort"
	"strings"
)

func moduleSymbols(
	all []*packages.Package,
	set *token.FileSet,
	modulePath string,
) []*module_symbol.Symbol {
	seen := make(map[string]*module_symbol.Symbol)

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
				module_symbol.New(
					o.Pkg().Path(),
					selectionOwner(selection),
					o.Name(),
					o,
					set.Position(expression.Pos()),
				),
			)
		}

		for identity, o := range loaded.TypesInfo.Uses {
			if !objectBelongsToModule(o, modulePath) ||
				o.Parent() != o.Pkg().Scope() {
				continue
			}

			addModuleSymbol(
				seen,
				module_symbol.New(
					o.Pkg().Path(),
					"",
					o.Name(),
					o,
					set.Position(identity.Pos()),
				),
			)
		}
	}

	var keys []string

	for key := range seen {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	var result []*module_symbol.Symbol

	for _, key := range keys {
		result = append(result, seen[key])
	}

	return result
}
