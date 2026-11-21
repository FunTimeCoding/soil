package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"strings"
)

func kernel(source string) bool {
	return source == constant.KernelSource ||
		strings.HasPrefix(source, constant.KernelPrefix)
}
