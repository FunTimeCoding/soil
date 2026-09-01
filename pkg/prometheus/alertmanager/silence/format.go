package silence

import (
	"fmt"
	"github.com/docker/go-units"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/time/constant"
)

func (s *Silence) Format(f *option.Format) string {
	t := status.New(f).String(
		s.formatRule(f),
		s.Author,
		fmt.Sprintf("%s ago", units.HumanDuration(s.Age())),
	).RawList(s)

	if r := s.Remain(); r > 0 {
		t.String(fmt.Sprintf("%s remain", units.HumanDuration(r)))
	}

	t.String(s.End.Format(constant.DateMinute))

	if f.HasTag(consoleConstant.TagState) {
		t.String(s.State)
	}

	t.DetailLink(s.Link, "Silence", "")

	if v := s.formatComment(f); v != "" {
		t.TagLine(consoleConstant.TagComment, "  %s", v)
	}

	return t.Format()
}
