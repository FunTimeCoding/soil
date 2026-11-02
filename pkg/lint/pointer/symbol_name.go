package pointer

import (
	"go/token"
	"strings"
)

func SymbolName(
	reference string,
	directory string,
) string {
	rest, found := strings.CutPrefix(Normalize(reference), directory)

	if !found || rest == "" {
		return ""
	}

	rest = rest[1:]

	if i := strings.IndexAny(rest, ".("); i != -1 {
		rest = rest[:i]
	}

	if !token.IsIdentifier(rest) {
		return ""
	}

	return rest
}
