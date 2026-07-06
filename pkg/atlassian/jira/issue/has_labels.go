package issue

import "github.com/funtimecoding/soil/pkg/strings/contains"

func (i *Issue) HasLabels(s ...string) bool {
	return contains.All(i.Raw.Fields.Labels, s)
}
