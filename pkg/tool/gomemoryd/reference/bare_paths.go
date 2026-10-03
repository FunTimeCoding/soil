package reference

import (
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"slices"
	"strings"
)

func barePaths(
	text string,
	roots []string,
) []*Finding {
	var result []*Finding

	for _, field := range strings.Fields(text) {
		word := strings.TrimRight(
			strings.TrimSuffix(
				strings.TrimRight(
					strings.TrimLeft(field, constant.ProseOpening),
					constant.ProseClosing,
				),
				constant.Possessive,
			),
			constant.ProseClosing,
		)

		if strings.Contains(word, constant.LocatorSeparator) {
			continue
		}

		first, _, found := strings.Cut(word, stringsConstant.Slash)

		if !found {
			continue
		}

		if first == libraryConstant.ParentDirectory ||
			slices.Contains(roots, first) {
			result = append(result, NewFinding(word, constant.BarePathText))
		}
	}

	return result
}
