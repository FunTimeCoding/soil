package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"strconv"
)

func backports(v string) (string, int) {
	m := constant.BackportsPattern.FindStringSubmatchIndex(v)

	if m == nil {
		return v, 0
	}

	n, e := strconv.Atoi(v[m[2]:m[3]])

	if e != nil {
		return v, 0
	}

	return v[:m[0]], n
}
