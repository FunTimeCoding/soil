package go_mod

import "github.com/funtimecoding/soil/pkg/go_mod/constant"

func IsDeadTag(stderr string) bool {
	return constant.DeadTagPattern.MatchString(stderr)
}
