package issue

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) Format(f *option.Format) string {
	s := status.New(f).String(
		i.Project,
		i.Title,
		i.formatType(f),
		i.formatAge(f),
	)
	s.DetailLink(i.Link, "Sentry", "")

	return s.Format()
}
