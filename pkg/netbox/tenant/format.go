package tenant

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Tenant) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(constant.TagIdentifier) {
		s.Integer32(t.Identifier)
	}

	s.String(t.formatName(f), t.Group).RawList(t.Raw)

	return s.Format()
}
