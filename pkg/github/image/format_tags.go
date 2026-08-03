package image

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (i *Image) formatTags(f *option.Format) string {
	if len(i.Tags) == 0 {
		if f.UseColor {
			return consoleConstant.Yellow("%s", constant.NoTags)
		}

		return constant.NoTags
	}

	return join.CommaSpace(i.Tags)
}
