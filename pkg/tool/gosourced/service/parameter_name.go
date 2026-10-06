package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/token"
	"unicode"
	"unicode/utf8"
)

func parameterName(field string) string {
	first, size := utf8.DecodeRuneInString(field)
	result := join.Empty(string(unicode.ToLower(first)), field[size:])

	if token.IsKeyword(result) || result == "new" {
		return join.Empty(result, "Value")
	}

	return result
}
