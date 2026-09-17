package pointer

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
	"unicode"
)

func PackageDirectory(s string) string {
	trimmed := strings.TrimSuffix(Normalize(s), "()")
	directory := ""
	base := trimmed

	if i := strings.LastIndex(trimmed, constant.Slash); i != -1 {
		directory = trimmed[:i]
		base = trimmed[i+1:]
	}

	if base == "" || unicode.IsUpper(rune(base[0])) {
		return directory
	}

	if j := strings.Index(base, constant.Dot); j != -1 {
		base = base[:j]
	}

	if directory == "" {
		return base
	}

	return join.Empty(directory, constant.Slash, base)
}
