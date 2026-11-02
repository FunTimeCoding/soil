package region

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Replace(
	content string,
	name string,
	block string,
) (string, error) {
	inner, closing, e := locate(content, name)

	if e != nil {
		return "", e
	}

	return join.Empty(
		content[:inner],
		stringsConstant.Unix,
		block,
		stringsConstant.Unix,
		content[closing:],
	), nil
}
