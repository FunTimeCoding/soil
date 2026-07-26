package session

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (s *Session) Format(f *option.Format) string {
	t := status.New(f)

	if f.HasTag(constant.TagIdentifier) {
		t.String(s.Identifier)
	}

	t.String(s.formatTabTitle(f), s.formatJob(f))

	return t.Format()
}
