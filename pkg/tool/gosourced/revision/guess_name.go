package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"strings"
)

func guessName(importPath string) string {
	elements := strings.Split(importPath, "/")
	result := elements[len(elements)-1]

	if isVersionElement(result) && len(elements) > 1 {
		result = elements[len(elements)-2]
	}

	result = strings.TrimPrefix(result, constant.ToolPrefix)
	result = strings.TrimSuffix(result, constant.ToolSuffix)
	result, _, _ = strings.Cut(result, constant.DotName)

	return result
}
