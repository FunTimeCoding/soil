package override

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/compact"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time"
)

func (o *Override) Format(f *option.Format) string {
	s := status.New(f).String(
		compact.Mail(o.User),
		o.Start.Format(time.DateMinute),
		o.End.Format(time.DateMinute),
	).RawList(o.Raw)

	if r := o.rotations(); len(r) > 0 {
		s.String(join.Comma(r))
	}

	return s.Format()
}
