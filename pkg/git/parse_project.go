package git

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/git/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func ParseProject(path string) (string, string) {
	parts := strings.Split(
		strings.Trim(path, stringsConstant.Slash),
		stringsConstant.Slash,
	)
	count := len(parts)

	if count != 2 {
		unexpected.Integer(count)
	}

	namespace := parts[0]
	repository := strings.TrimSuffix(parts[count-1], constant.Directory)

	return namespace, repository
}
