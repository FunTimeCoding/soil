package installed

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"path"
	"strings"
)

func Touches(
	changed []string,
	directories map[string]bool,
) bool {
	for _, f := range changed {
		if !strings.HasSuffix(f, constant.TestSuffix) &&
			directories[path.Dir(f)] {
			return true
		}
	}

	return false
}
