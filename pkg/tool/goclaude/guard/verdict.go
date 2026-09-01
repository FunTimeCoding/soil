package guard

import "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"

func Verdict(
	system string,
	command string,
) string {
	npx, e := localNpx(command)

	if e != nil && constant.NpxInvocation.MatchString(command) {
		return constant.NpxMessage
	}

	if npx {
		return constant.NpxMessage
	}

	pip, f := localPipInstall(command)

	if f != nil && constant.PipInvocation.MatchString(command) {
		return constant.PipMessage
	}

	if pip {
		return constant.PipMessage
	}

	if system != "darwin" {
		return ""
	}

	inPlace, g := localSedInPlace(command)

	if g != nil && constant.SedInvocation.MatchString(command) {
		return constant.SedMessage
	}

	if inPlace {
		return constant.SedMessage
	}

	return ""
}
