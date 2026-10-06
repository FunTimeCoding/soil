package format

import (
	"github.com/funtimecoding/soil/pkg/integers64"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func AlwaysMemory(m *record.Memory) string {
	block := join.Empty(
		constant.MemoryHeadingPrefix,
		m.Name,
		constant.IdentifierOpen,
		integers64.ToString(m.Identifier),
		constant.IdentifierClose,
		tagSuffix(m.Tags),
		stringConstant.Unix,
		baseLine(m.Metadata),
		m.Content,
		stringConstant.Unix,
	)

	if len(m.Children) == 0 {
		return block
	}

	return join.Empty(
		block,
		constant.ChildIndent,
		join.CommaSpace(m.Children),
		stringConstant.Unix,
	)
}
