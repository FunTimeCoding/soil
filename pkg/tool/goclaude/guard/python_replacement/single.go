package python_replacement

import "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"

func single(code string) (string, bool) {
	if len(constant.PythonReplaceCall.FindAllString(code, -1)) != 1 {
		return "", false
	}

	if constant.PythonReplaceCount.MatchString(code) ||
		constant.PythonRegexModule.MatchString(code) ||
		constant.PythonLoop.MatchString(code) {
		return "", false
	}

	if len(constant.PythonRead.FindAllString(code, -1)) != 1 ||
		len(constant.PythonWrite.FindAllString(code, -1)) != 1 {
		return "", false
	}

	tokens := map[string]bool{}
	var token string

	for _, m := range constant.PythonFileTarget.FindAllStringSubmatch(
		code,
		-1,
	) {
		tokens[m[1]] = true
		token = m[1]
	}

	if len(tokens) != 1 {
		return "", false
	}

	return target(code, token), true
}
