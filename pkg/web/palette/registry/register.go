package registry

import "github.com/funtimecoding/soil/pkg/web/palette"

func (r *Registry) Register(commands ...*palette.Command) {
	r.commands = append(r.commands, commands...)
}
