package image

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (i *Image) formatTags(f *option.Format) string {
	if len(i.Tags) == 0 {
		if f.UseColor {
			return console.Yellow("%s", NoTags)
		}

		return NoTags
	}

	return join.CommaSpace(i.Tags)
}
