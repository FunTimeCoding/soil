package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
)

func (p *Page) Commands() []*palette.Command {
	if !p.HasCommandPalette() {
		return nil
	}

	var result []*palette.Command

	if name := p.identity.Name(); name != "" {
		result = append(
			result,
			palette.NewCommand(
				name,
				constant.RootPath,
				constant.PaletteNavigate,
			),
		)
	}

	for _, i := range p.items {
		result = append(
			result,
			palette.NewCommand(i.Label(), i.Path(), constant.PaletteNavigate),
		)
	}

	return append(result, p.commands...)
}
