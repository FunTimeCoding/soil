package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module_symbol"
)

func moduleMemberName(symbol *module_symbol.Symbol) string {
	if symbol.Owner == "" {
		return symbol.Name
	}

	return join.Empty(symbol.Owner, constant.MemberSeparator, symbol.Name)
}
