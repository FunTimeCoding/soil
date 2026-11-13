package gobreak

import (
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"strings"
)

func compilerLine(output []string) string {
	for _, l := range output {
		if !strings.HasPrefix(l, constant.PackageHeader) {
			return l
		}
	}

	return output[0]
}
