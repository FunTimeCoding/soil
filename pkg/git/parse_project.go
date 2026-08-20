package git

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/git/constant"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
)

func ParseProject(path string) (string, string) {
	parts := split.Slash(strings.Trim(path, stringConstant.Slash))
	count := len(parts)

	if count != 2 {
		unexpected.Integer(count)
	}

	namespace := parts[0]
	repository := strings.TrimSuffix(parts[count-1], constant.Directory)

	return namespace, repository
}
