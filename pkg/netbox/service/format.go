package service

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (s *Service) Format(f *option.Format) string {
	t := status.New(f)

	if f.HasTag(tag.Identifier) {
		t.Integer32(s.Identifier)
	}

	t.String(s.formatName(f)).RawList(s.Raw)

	return t.Format()
}
