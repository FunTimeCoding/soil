package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module_symbol"
	"go/types"
)

func moduleMember(
	target *types.Package,
	symbol *module_symbol.Symbol,
) types.Object {
	if symbol.Owner == "" {
		return target.Scope().Lookup(symbol.Name)
	}

	owner := target.Scope().Lookup(symbol.Owner)

	if owner == nil {
		return nil
	}

	found, _, _ := types.LookupFieldOrMethod(
		owner.Type(),
		true,
		target,
		symbol.Name,
	)

	return found
}
