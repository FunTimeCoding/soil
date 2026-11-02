package profile

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func sectionText(
	heading string,
	body []string,
) string {
	return join.Empty(
		heading,
		stringsConstant.Unix,
		stringsConstant.Unix,
		join.NewLine(body),
	)
}
