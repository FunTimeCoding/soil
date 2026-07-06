package repository

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/strings/split"
)

func (r *Repository) Format(f *option.Format) string {
	s := status.New(f).String(
		r.formatName(f),
		r.Path,
		r.formatConcern(f),
	).RawList(r)

	if !r.IsClean && r.Status != "" {
		for _, l := range split.NewLine(r.Status) {
			s.TagLine(tag.Changes, "%s", l)
		}
	}

	return s.Format()
}
