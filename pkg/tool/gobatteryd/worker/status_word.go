package worker

import "github.com/funtimecoding/soil/pkg/tool/gobatteryd/constant"

func statusWord(flag string) string {
	if word, found := constant.StatusWords[flag]; found {
		return word
	}

	return flag
}
