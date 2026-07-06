package rack_type

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (t *Type) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(t.Identifier)
	}

	s.String(t.formatName(f)).RawList(t.Raw)

	return s.Format()
}
