package page

import "github.com/funtimecoding/soil/pkg/atlassian/constant"

func ToStorage(markdown string) string {
	if constant.MacroCommentPattern.MatchString(markdown) {
		return markersToMacros(markdown)
	}

	return markdownToMarkup(markdown)
}
