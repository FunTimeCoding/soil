package goflow

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goflow/constant"
	"slices"
	"strings"
)

func front(source []byte) int {
	lines := strings.Split(string(source), "\n")

	if len(lines) == 0 || lines[0] != constant.Delimiter {
		return 0
	}

	at := slices.Index(lines[1:], constant.Delimiter)

	if at < 0 {
		return 0
	}

	return len(join.NewLine(lines[:at+2])) + 1
}
