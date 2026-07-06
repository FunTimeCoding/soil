package gofix

import (
	"github.com/funtimecoding/soil/pkg/tool/gofix/constant"
	"go/ast"
	"go/token"
)

func collapseSingleParameter(
	fileSet *token.FileSet,
	f *ast.FuncDecl,
	source []byte,
) []edit {
	params := f.Type.Params

	if params == nil || len(params.List) != 1 {
		return nil
	}

	openPosition := fileSet.Position(params.Opening)
	closePosition := fileSet.Position(params.Closing)

	if openPosition.Line == closePosition.Line {
		return nil
	}

	field := params.List[0]
	parameterStart := fileSet.Position(field.Pos()).Offset
	parameterEnd := fileSet.Position(field.End()).Offset
	parameterText := string(source[parameterStart:parameterEnd])
	lineStart := openPosition.Offset - (openPosition.Column - 1)
	prefix := source[lineStart:openPosition.Offset]
	closeOffset := closePosition.Offset
	suffixEnd := closeOffset

	for suffixEnd < len(source) && source[suffixEnd] != '\n' {
		suffixEnd++
	}

	suffix := source[closeOffset:suffixEnd]
	collapsedLength := len(prefix) + 1 + len(parameterText) + len(suffix)

	if collapsedLength > constant.MaxSingleParameterLength {
		return nil
	}

	return []edit{
		{
			position: params.Opening + 1,
			end:      params.Closing,
			newText:  parameterText,
		},
	}
}
