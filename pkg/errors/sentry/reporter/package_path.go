package reporter

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func packagePath(function string) string {
	slash := strings.LastIndex(function, stringsConstant.Slash)

	if i := strings.Index(function[slash+1:], stringsConstant.Dot); i >= 0 {
		return function[:slash+1+i]
	}

	return function
}
