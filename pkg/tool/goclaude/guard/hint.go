package guard

import "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"

func Hint(command string) string {
	if localOapiCodegen(command) {
		return constant.OapiCodegenHint
	}

	return ""
}
