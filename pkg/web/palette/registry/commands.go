package registry

import "github.com/funtimecoding/soil/pkg/web/palette"

func (r *Registry) Commands() []*palette.Command {
	return r.commands
}
