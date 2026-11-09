package pointer

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/token"
	"strings"
)

func MemberName(
	reference string,
	directory string,
) string {
	head := SymbolName(reference, directory)

	if head == "" {
		return ""
	}

	rest, _ := strings.CutPrefix(Normalize(reference), directory)
	rest, found := strings.CutPrefix(
		rest[1:],
		join.Empty(head, stringsConstant.Dot),
	)

	if !found {
		return ""
	}

	if i := strings.IndexAny(rest, ".("); i != -1 {
		rest = rest[:i]
	}

	if !token.IsIdentifier(rest) {
		return ""
	}

	return rest
}
