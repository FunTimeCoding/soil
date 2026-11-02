package image

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Image) formatDigest(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", i.Digest)
	}

	return i.Digest
}
