package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func withRewritten(
	text string,
	rewritten []*record.Memory,
) string {
	if len(rewritten) == 0 {
		return text
	}

	var names []string

	for _, m := range rewritten {
		names = append(names, fmt.Sprintf("%d %s", m.Identifier, m.Name))
	}

	return join.NewLine(
		[]string{
			text,
			join.Empty(constant.RewrittenPrefix, join.CommaSpace(names)),
		},
	)
}
