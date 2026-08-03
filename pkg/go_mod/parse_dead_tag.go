package go_mod

import "github.com/funtimecoding/soil/pkg/go_mod/constant"

func ParseDeadTag(stderr string) (string, string) {
	m := constant.DeadTagPattern.FindStringSubmatch(stderr)

	if m == nil {
		return "", ""
	}

	return m[1], m[2]
}
