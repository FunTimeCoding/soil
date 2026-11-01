package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"strconv"
	"strings"
)

func isVersionElement(element string) bool {
	number, found := strings.CutPrefix(element, constant.VersionPrefix)

	if !found {
		return false
	}

	_, e := strconv.Atoi(number)

	return e == nil
}
