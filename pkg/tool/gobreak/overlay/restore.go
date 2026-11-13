package overlay

import (
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
)

func (o *Overlay) Restore(
	root string,
	text string,
) string {
	for original, target := range o.Replace {
		name := system.RelativePath(root, original)
		text = strings.ReplaceAll(text, target, name)
		text = strings.ReplaceAll(text, system.RelativePath(root, target), name)
	}

	return text
}
