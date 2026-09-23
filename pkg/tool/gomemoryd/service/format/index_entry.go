package format

import (
	"github.com/funtimecoding/soil/pkg/integers64"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func IndexEntry(m *store.MemorySummary) string {
	line := join.Empty(
		integers64.ToString(m.Identifier),
		stringConstant.Space,
		m.Name,
		tagSuffix(m.Tags),
		constant.EntrySeparator,
		m.Description,
	)

	if len(m.Children) == 0 {
		return line
	}

	return join.Empty(
		line,
		stringConstant.Unix,
		constant.ChildIndent,
		join.CommaSpace(m.Children),
	)
}
