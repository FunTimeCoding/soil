package manufacturer

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (m *Manufacturer) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(consoleConstant.TagIdentifier) {
		s.Integer32(m.Identifier)
	}

	s.String(m.formatName(f)).RawList(m.Raw)

	return s.Format()
}
