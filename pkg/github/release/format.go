package release

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
)

func (r *Release) Format(f *option.Format) string {
	return status.New(f).String(
		r.formatName(f),
		r.Create.Format(timeConstant.DateMinute),
	).RawList(r.Raw).Format()
}
