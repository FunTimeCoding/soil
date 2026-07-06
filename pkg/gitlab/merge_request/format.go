package merge_request

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Request) Format(f *option.Format) string {
	s := status.New(f).Integer64(
		r.Project,
		r.Identifier,
	).String(
		r.formatState(f),
		r.formatTitle(f),
		r.formatAge(f),
	).RawList(r.Raw)
	s.DetailLink(r.Link, "GitLab", "")

	return s.Format()
}
