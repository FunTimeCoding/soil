package rack

import "github.com/funtimecoding/soil/pkg/netbox/tag"

func (r *Rack) tags() []string {
	return tag.Names(r.Raw.Tags)
}
