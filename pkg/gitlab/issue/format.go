package issue

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) Format(f *option.Format) string {
	return status.New(f).Integer64(i.Identifier).String(
		i.State,
		i.Title,
	).RawList(i.Raw).Format()
}
