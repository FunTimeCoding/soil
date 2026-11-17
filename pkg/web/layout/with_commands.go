package layout

import "github.com/funtimecoding/soil/pkg/web/palette"

func (p *Page) WithCommands(commands func() []*palette.Command) *Page {
	p.commands = commands

	return p
}
