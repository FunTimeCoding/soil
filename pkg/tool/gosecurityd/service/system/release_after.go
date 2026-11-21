package system

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"

func releaseAfter(number int) string {
	for name, n := range constant.ReleaseNumbers {
		if n == number+1 {
			return name
		}
	}

	return ""
}
