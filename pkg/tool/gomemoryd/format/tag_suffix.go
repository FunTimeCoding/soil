package format

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func tagSuffix(tags []string) string {
	if len(tags) == 0 {
		return ""
	}

	return join.Empty(constant.TagOpen, join.Space(tags...), constant.TagClose)
}
