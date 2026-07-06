package image

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Image) formatDigest(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", i.Digest)
	}

	return i.Digest
}
