package region

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
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
		constant.Unix,
		block,
		constant.Unix,
		content[closing:],
	), nil
}
