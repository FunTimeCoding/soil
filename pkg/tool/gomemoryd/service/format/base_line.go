package format

import (
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"
)

func baseLine(metadata map[string]string) string {
	bases := reference.Bases(metadata[constant.BaseKey])

	if len(bases) == 0 {
		return ""
	}

	var quoted []string

	for _, base := range bases {
		quoted = append(
			quoted,
			join.Empty(constant.Backtick, base, constant.Backtick),
		)
	}

	return join.Empty(
		constant.BasePrefix,
		join.CommaSpace(quoted),
		stringConstant.Unix,
	)
}
