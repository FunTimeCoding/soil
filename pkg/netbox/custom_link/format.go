package custom_link

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (l *Link) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(l.Identifier)
	}

	s.String(l.formatName(f)).RawList(l.Raw)

	return s.Format()
}
