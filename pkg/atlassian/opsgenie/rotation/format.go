package rotation

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/compact"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Rotation) Format(f *option.Format) string {
	s := status.New(f).String(
		r.Name,
		string(r.Type),
		r.Start.Format("Monday"),
		r.current(),
	).RawList(r.Raw)

	for _, participant := range r.Participants {
		s.Line("    Member: %s", compact.Mail(participant.Username))
	}

	return s.Format()
}
