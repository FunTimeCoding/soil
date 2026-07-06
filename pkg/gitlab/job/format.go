package job

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (j *Job) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer64(j.Identifier)
	}

	s.String(j.formatName(f))

	if f.HasTag(tag.Project) {
		s.String(j.formatProject())
	}

	s.String(
		j.formatUser(),
		j.formatDate(f),
		j.formatConcern(f),
	)
	s.DetailLink(j.Link, "GitLab", "")

	return s.RawList(j.Raw).Format()
}
