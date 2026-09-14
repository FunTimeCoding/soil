package service

import "github.com/funtimecoding/soil/pkg/lint/concern"

func moduleBreakage(
	symbol *ModuleSymbol,
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
