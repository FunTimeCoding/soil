package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"slices"
)

func (w *Walk) invoke(arguments []*syntax.Word) bool {
	words := literalWords(arguments)

	if len(words) == 0 {
		return w.block(constant.BlockerExpansion)
	}

	if escapes(words[0], arguments[1:]) {
		return w.block(w.commandLabel(words))
	}

	if slices.Contains(constant.Wrappers, words[0]) {
		return w.unwrap(words[0], arguments[1:])
	}

	if !slices.Contains(w.classes, words[0]) {
		w.classes = append(w.classes, words[0])
	}

	switch words[0] {
	case constant.ChangeDirectory:
		return w.changeDirectory(arguments[1:])
	case constant.GitCommand:
		if !w.allowlisted(words) {
			return w.block(w.commandLabel(words))
		}

		if w.moved {
			return w.block(constant.BlockerGitAfterChangeDirectory)
		}

		return w.source(true, constant.AllowlistSource)
	case constant.RemoveCommand, constant.MakeDirectory:
		return w.source(
			w.scratchOperands(arguments[1:], 1),
			constant.ScratchSource,
		) || w.block(w.commandLabel(words))
	case constant.MoveCommand:
		return w.source(
			w.scratchOperands(arguments[1:], 2),
			constant.ScratchSource,
		) || w.block(w.commandLabel(words))
	case constant.CopyCommand:
		return w.source(
			w.copyOperands(arguments[1:]),
			constant.ScratchSource,
		) || w.block(w.commandLabel(words))
	case constant.ReadCommand:
		w.forgetOperands(arguments[1:])
	case constant.PrintfCommand:
		w.forgetPrintfTarget(arguments[1:])
	}

	if slices.Contains(constant.ImplicitCommands, words[0]) {
		return w.source(true, constant.BuiltInSource)
	}

	return w.source(w.allowlisted(words), constant.AllowlistSource) ||
		w.block(w.commandLabel(words))
}
