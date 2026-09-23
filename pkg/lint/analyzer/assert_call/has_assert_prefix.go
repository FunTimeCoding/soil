package assert_call

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func HasAssertPrefix(name string) bool {
	return strings.HasPrefix(strings.ToLower(name), constant.AssertHelperPrefix)
}
