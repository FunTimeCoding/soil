package event

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/time"
)

func (e *Event) Format(f *option.Format) string {
	s := status.New(f)
	s.String(e.Reason)
	s.String(e.Type)
	s.String(e.RegardingKind)
	s.String(e.Namespace)
	s.String(e.formatAge())

	if f.HasTag(tag.Cluster) {
		s.String("  Cluster: %s", e.Cluster)
	}

	s.Line("  Note: %s", e.Note)
	s.Line("  Create: %s", e.Create.Format(time.DateMinute))
	s.RawList(e.Raw)

	return s.Format()
}
