package post

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/text"
	"github.com/funtimecoding/soil/pkg/time"
)

func (p *Post) Format(f *option.Format) string {
	lines := split.NewLine(
		text.OptimizeWhitespace(p.Message, constant.CompactWhitespace),
	)
	s := status.New(f).String(
		p.Create.Format(time.HourMinute),
		p.formatUser(),
		formatLine(f, lines[0]),
	).RawList(p.Raw)

	if l := len(lines); l > 1 {
		for i := 1; i < l; i++ {
			s.Line("%s", formatLine(f, lines[i]))
		}
	}

	return s.Format()
}
