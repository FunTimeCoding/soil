package layout

import "github.com/funtimecoding/soil/pkg/web/palette"

func (p *Page) WithCommands(commands ...*palette.Command) *Page {
	p.commands = append(p.commands, commands...)

	return p
}
