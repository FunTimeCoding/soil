package component

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/url"
)

func RowPath(
	page string,
	action string,
	parameter string,
	name string,
) string {
	return join.Empty(
		page,
		stringsConstant.Slash,
		action,
		stringsConstant.Question,
		url.Values{parameter: {name}}.Encode(),
	)
}
