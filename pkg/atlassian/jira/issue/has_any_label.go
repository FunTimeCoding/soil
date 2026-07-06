package issue

import "github.com/funtimecoding/soil/pkg/strings/contains"

func (i *Issue) HasAnyLabel(s ...string) bool {
	return contains.Any(i.Raw.Fields.Labels, s)
}
