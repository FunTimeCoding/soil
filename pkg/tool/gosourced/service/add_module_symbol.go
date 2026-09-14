package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func addModuleSymbol(
	seen map[string]*ModuleSymbol,
	symbol *ModuleSymbol,
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
