package reporter

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func packagePath(function string) string {
	slash := strings.LastIndex(function, constant.Slash)

	if i := strings.Index(function[slash+1:], constant.Dot); i >= 0 {
		return function[:slash+1+i]
	}

	return function
}
