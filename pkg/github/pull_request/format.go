package pull_request

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/time"
)

func (r *Request) Format(f *option.Format) string {
	return status.New(f).String(
		r.formatName(f),
		r.Create.Format(time.DateMinute),
	).RawList(r.Raw).Format()
}
