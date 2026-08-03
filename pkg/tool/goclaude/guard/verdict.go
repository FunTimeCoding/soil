package guard

import "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"

func Verdict(
	system string,
	command string,
) string {
	if system != "darwin" {
		return ""
	}

	inPlace, e := localSedInPlace(command)

	if e != nil && constant.SedInvocation.MatchString(command) {
		return constant.SedMessage
	}

	if inPlace {
		return constant.SedMessage
	}

	return ""
}
