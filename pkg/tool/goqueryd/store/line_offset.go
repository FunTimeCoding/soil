package store

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func LineOffset(
	body string,
	line int,
) int {
	result := 0

	for range line - 1 {
		next := strings.Index(body[result:], constant.Unix)

		if next < 0 {
			return result
		}

		result += next + 1
	}

	return result
}
