package virtual_machine

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Machine) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(consoleConstant.TagIdentifier) {
		s.Integer32(r.Identifier)
	}

	s.String(r.formatName(f)).RawList(r.Raw)

	return s.Format()
}
