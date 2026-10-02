package dependency

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (d *Dependency) Format(f *option.Format) string {
	return status.New(f).String(
		d.formatPath(f),
		d.Version,
		d.formatFamily(f),
		join.Comma(d.Identifiers),
		d.formatConcern(f),
	).RawList(
		d,
	).Format()
}
