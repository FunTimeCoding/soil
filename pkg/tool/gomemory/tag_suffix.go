package gomemory

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	gomemorydConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func tagSuffix(tags *[]string) string {
	if tags == nil || len(*tags) == 0 {
		return ""
	}

	return join.Empty(
		gomemorydConstant.TagOpen,
		join.Space(*tags...),
		gomemorydConstant.TagClose,
	)
}
