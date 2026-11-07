package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
)

func (p *Page) Commands() []*palette.Command {
	var result []*palette.Command

	for _, i := range p.items {
		result = append(
			result,
			palette.NewCommand(i.Label(), i.Path(), constant.PaletteNavigate),
		)
	}

	return append(result, p.commands...)
}
