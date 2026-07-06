package issue

import "github.com/funtimecoding/soil/pkg/console/description"

func (i *Issue) Meta() *description.Description {
	return description.New("Issue", "Issue")
}
