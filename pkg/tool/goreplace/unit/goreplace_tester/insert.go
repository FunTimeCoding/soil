package goreplace_tester

import (
	textConstant "github.com/funtimecoding/soil/pkg/text/constant"
	"strings"
)

func Insert(
	opening string,
	anchor string,
	insert string,
) string {
	return strings.Join(
		[]string{
			opening,
			anchor,
			textConstant.DividerMarker,
			insert,
			textConstant.ReplaceMarker,
			"",
		},
		"\n",
	)
}
