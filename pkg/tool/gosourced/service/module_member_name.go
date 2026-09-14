package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func moduleMemberName(symbol *ModuleSymbol) string {
	if symbol.Owner == "" {
		return symbol.Name
	}

	return join.Empty(symbol.Owner, constant.MemberSeparator, symbol.Name)
}
