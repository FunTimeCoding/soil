package reach

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"mvdan.cc/sh/v3/syntax"
	"path/filepath"
	"slices"
)

func visitCall(
	call *syntax.CallExpr,
	directory string,
	result *CommandWrites,
) string {
	if len(call.Args) == 0 {
		return directory
	}

	name := filepath.Base(call.Args[0].Lit())
	flags := call.Args[1:]
	arguments := operands(flags)

	switch {
	case name == constant.ChangeDirectoryCommand && len(arguments) > 0:
		if v, okay := literalValue(arguments[0]); okay {
			return absolute(directory, v)
		}
	case name == constant.GitCommand && len(arguments) > 0 &&
		slices.Contains(constant.GitPathCommands, arguments[0].Lit()):
		addWords(result, directory, arguments[1:])
	case slices.Contains(constant.PathCommands, name):
		addWords(result, directory, arguments)
	case name == constant.CopyCommand && len(arguments) > 0:
		addWords(result, directory, arguments[len(arguments)-1:])
	case slices.Contains(constant.SedCommands, name) &&
		hasFlag(flags, constant.InPlaceFlag, constant.InPlaceLongFlag) &&
		len(arguments) > 1:
		addWords(result, directory, arguments[1:])
	case isFixer(name, flags):
		result.Fixer = true
	case slices.Contains(constant.ScriptCommands, name):
		result.Script = true
	}

	return directory
}
