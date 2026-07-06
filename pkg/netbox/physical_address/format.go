package physical_address

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Address) Format(f *option.Format) string {
	return status.New(f).String(
		a.formatName(f),
		a.formatDescription(),
		a.formatObjectType(f),
		a.formatInterface(),
		a.formatDevice(f),
	).RawList(a.Raw).Format()
}
