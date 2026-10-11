package module

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (m *Module) Format(f *option.Format) string {
	return status.New(f).String(
		m.formatPath(f),
		m.formatVersion(),
		m.formatRelease(),
		m.formatUntagged(),
		m.formatConcern(f),
	).RawList(
		m,
	).Format()
}
