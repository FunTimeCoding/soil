package module_bay

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (b *Bay) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(b.Identifier)
	}

	s.String(b.formatName(f), b.Description).RawList(b.Raw)

	return s.Format()
}
