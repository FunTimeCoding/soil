package refusal

import (
	"strconv"
	"strings"
)

func status(text string) int {
	code, _, _ := strings.Cut(text, " ")
	result, e := strconv.Atoi(code)

	if e != nil {
		return 0
	}

	return result
}
