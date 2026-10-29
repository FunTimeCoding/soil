package failure

import (
	"github.com/funtimecoding/soil/pkg/failure/constant"
	"strings"
)

func owner(name string) string {
	if i := strings.LastIndex(name, constant.NameSeparator); i >= 0 {
		return name[:i+1]
	}

	return name
}
