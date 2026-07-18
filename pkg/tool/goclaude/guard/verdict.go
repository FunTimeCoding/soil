package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"regexp"
)

var sedInvocation = regexp.MustCompile(`(^|[|&;(\s])sed(\s|$)`)

func verdict(
	system string,
	command string,
) string {
	if system == "darwin" && sedInvocation.MatchString(command) {
		return constant.SedMessage
	}

	return ""
}
