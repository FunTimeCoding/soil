package run

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/time"
)

func (r *Run) Format(f *option.Format) string {
	s := status.New(f).String(
		r.formatName(f),
		r.Status,
	)

	if f.HasTag(tag.Timestamp) {
		s.String(r.Create.Format(time.DateMinute))
	}

	s.String(r.formatConcern(f)).RawList(r.Raw)

	return s.Format()
}
