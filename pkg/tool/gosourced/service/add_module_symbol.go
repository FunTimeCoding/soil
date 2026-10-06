package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module_symbol"
)

func addModuleSymbol(
	seen map[string]*module_symbol.Symbol,
	symbol *module_symbol.Symbol,
) {
	key := join.Empty(
		symbol.PackagePath,
		constant.MemberSeparator,
		symbol.Owner,
		constant.MemberSeparator,
		symbol.Name,
	)

	if _, okay := seen[key]; !okay {
		seen[key] = symbol
	}
}
