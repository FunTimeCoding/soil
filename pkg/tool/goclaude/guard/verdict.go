package guard

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
)

func Verdict(
	system string,
	command string,
) string {
	calls, d := gitCalls(command)

	if d != nil && constant.GitInvocation.MatchString(command) {
		return fmt.Sprintf(constant.GitUnreadableMessage, "command")
	}

	for _, c := range calls {
		if b := c.Block(); b != "" {
			return b
		}
	}

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

	if localSingleReplacement(command) {
		return constant.SingleReplacementMessage
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

	unsupported, h := localXargsUnsupportedFlag(command)

	if h != nil && constant.XargsInvocation.MatchString(command) {
		return constant.XargsMessage
	}

	if unsupported {
		return constant.XargsMessage
	}

	return ""
}
