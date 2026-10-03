package mark

import "github.com/funtimecoding/soil/pkg/generative/constant"

func parseLocator(locator string) string {
	l := len(locator)
	prefix := len(constant.MarkUserPrefix)
	suffix := len(constant.MarkProfileSuffix)

	if l > prefix+suffix {
		return locator[prefix : l-suffix]
	}

	return ""
}
