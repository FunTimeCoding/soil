package commit

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
)

func (c *Commit) Format(f *option.Format) string {
	return status.New(f).String(
		c.Date.Format(timeConstant.DateMinute),
		c.Author,
		c.formatTitle(f),
	).RawList(c.Raw).Format()
}
