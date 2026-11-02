package repository

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/split"
)

func (r *Repository) Format(f *option.Format) string {
	s := status.New(f).String(
		r.formatName(f),
		r.Path,
		r.formatConcern(f),
	).RawList(
		r,
	)

	if !r.IsClean && r.Status != "" {
		for _, l := range split.NewLine(r.Status) {
			s.TagLine(consoleConstant.TagChanges, "%s", l)
		}
	}

	return s.Format()
}
