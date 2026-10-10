package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"strings"
)

func insertBlock(
	opening string,
	anchor string,
	insert string,
) string {
	return strings.Join(
		[]string{
			opening,
			anchor,
			constant.DividerMarker,
			insert,
			constant.ReplaceMarker,
			"",
		},
		"\n",
	)
}
