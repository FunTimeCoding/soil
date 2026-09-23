package format

import (
	"github.com/funtimecoding/soil/pkg/floats"
	"github.com/funtimecoding/soil/pkg/integers64"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func RelevantMemory(r *store.SearchResult) string {
	return join.Empty(
		constant.MemoryHeadingPrefix,
		r.Name,
		constant.IdentifierOpen,
		integers64.ToString(r.Identifier),
		constant.IdentifierClose,
		tagSuffix(r.Tags),
		constant.RankSeparator,
		floats.ToStringRounded(r.Rank),
		stringConstant.Unix,
		r.Content,
		stringConstant.Unix,
	)
}
