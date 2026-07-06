package project

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (p *Project) Format(f *option.Format) string {
	s := status.New(f).Integer64(
		p.Identifier,
	).String(p.formatName(f)).RawList(p.Raw)
	s.DetailLink(p.Link, "GitLab", "")

	return s.Format()
}
