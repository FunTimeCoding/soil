package gofix

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gofix/constant"
	"go/format"
	"strings"
)

func joinParameter(text string) (string, bool) {
	formatted, e := format.Source(
		[]byte(
			join.Empty(
				constant.ParameterSourcePrefix,
				strings.ReplaceAll(text, "\n", " "),
				constant.ParameterSourceSuffix,
			),
		),
	)

	if e != nil {
		return "", false
	}

	result, found := strings.CutPrefix(
		string(formatted),
		constant.ParameterSourcePrefix,
	)

	if !found {
		return "", false
	}

	result, found = strings.CutSuffix(result, constant.ParameterSourceSuffix)

	if !found || strings.Contains(result, "\n") {
		return "", false
	}

	return result, true
}
