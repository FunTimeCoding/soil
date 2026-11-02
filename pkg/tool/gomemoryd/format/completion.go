package format

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
)

func Completion(
	name string,
	body string,
) string {
	return join.Empty(
		constant.MemoryHeadingPrefix,
		name,
		stringsConstant.Unix,
		body,
		stringsConstant.Unix,
	)
}
