package printer

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/face"
)

func (p *Printer) Print(
	f face.Formattable,
	indent int,
) {
	var title string

	if p.format.UseCompact {
		title = f.Meta().Short
	} else {
		title = f.Meta().Title
	}

	console.Print(f, title, indent, p.format)
}
