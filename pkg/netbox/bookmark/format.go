package bookmark

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (b *Bookmark) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(consoleConstant.TagIdentifier) {
		s.Integer32(b.Identifier)
	}

	s.String(b.Display).String(b.ObjectType).String(b.Link).RawList(b.Raw)

	return s.Format()
}
