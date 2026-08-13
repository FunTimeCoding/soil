package run

import (
	console "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/time/constant"
)

func (r *Run) Format(f *option.Format) string {
	s := status.New(f).String(r.formatName(f), r.Status)

	if f.HasTag(console.TagTimestamp) {
		s.String(r.Create.Format(constant.DateMinute))
	}

	s.String(r.formatConcern(f)).RawList(r.Raw)

	return s.Format()
}
