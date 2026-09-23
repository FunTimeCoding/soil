package format

import (
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func Impression(i *store.Impression) string {
	return join.Empty(
		constant.MemoryHeadingPrefix,
		i.Source,
		stringConstant.Unix,
		i.Content,
		stringConstant.Unix,
	)
}
