package lint

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/repository"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path"
	"strings"
)

func outsideTrees(
	clean string,
	siblings []string,
) bool {
	if path.IsAbs(clean) || repository.InGitDirectory(clean) {
		return true
	}

	if !strings.HasPrefix(clean, soil.ParentDirectory) {
		return false
	}

	for _, s := range siblings {
		if clean == s ||
			strings.HasPrefix(clean, join.Empty(s, stringsConstant.Slash)) {
			return false
		}
	}

	return true
}
