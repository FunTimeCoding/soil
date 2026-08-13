package strings

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"regexp"
)

func ReduceSpace(s string) string {
	return regexp.MustCompile(`\s{2,}`).ReplaceAllString(s, constant.Space)
}
