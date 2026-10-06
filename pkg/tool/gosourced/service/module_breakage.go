package service

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module_symbol"
)

func moduleBreakage(
	symbol *module_symbol.Symbol,
	key string,
	text string,
) *concern.Concern {
	return concern.NewLine(
		key,
		text,
		symbol.Position.Filename,
		symbol.Position.Line,
		"",
		false,
	)
}
