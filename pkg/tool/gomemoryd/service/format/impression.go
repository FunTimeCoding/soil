package format

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
)

func Impression(i *record.Impression) string {
	return join.Empty(
		constant.MemoryHeadingPrefix,
		i.Source,
		stringsConstant.Unix,
		i.Content,
		stringsConstant.Unix,
	)
}
