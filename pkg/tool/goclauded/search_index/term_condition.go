package search_index

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
	"unicode/utf8"
)

func termCondition(term string) (string, string) {
	if utf8.RuneCountInString(term) >= 3 {
		return "block_text MATCH ?",
			join.Empty(`"`, strings.ReplaceAll(term, `"`, `""`), `"`)
	}

	escaped := strings.NewReplacer(`\`, `\\`, "%", `\%`, "_", `\_`).Replace(
		term,
	)

	return `block_text.body LIKE ? ESCAPE '\'`, join.Empty("%", escaped, "%")
}
