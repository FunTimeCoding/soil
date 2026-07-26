package site

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (s *Site) Format(f *option.Format) string {
	t := status.New(f)

	if f.HasTag(constant.TagIdentifier) {
		t.Integer32(s.Identifier)
	}

	t.String(s.formatName(f)).RawList(s.Raw)

	return t.Format()
}
