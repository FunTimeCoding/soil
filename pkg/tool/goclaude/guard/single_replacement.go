package guard

import "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"

func singleReplacement(source string) (string, bool) {
	if len(constant.PythonReplaceCall.FindAllString(source, -1)) != 1 {
		return "", false
	}

	if constant.PythonReplaceCount.MatchString(source) ||
		constant.PythonRegexModule.MatchString(source) ||
		constant.PythonLoop.MatchString(source) {
		return "", false
	}

	if len(constant.PythonRead.FindAllString(source, -1)) != 1 ||
		len(constant.PythonWrite.FindAllString(source, -1)) != 1 {
		return "", false
	}

	targets := map[string]bool{}
	var target string

	for _, m := range constant.PythonFileTarget.FindAllStringSubmatch(
		source,
		-1,
	) {
		targets[m[1]] = true
		target = m[1]
	}

	if len(targets) != 1 {
		return "", false
	}

	return pythonTarget(source, target), true
}
