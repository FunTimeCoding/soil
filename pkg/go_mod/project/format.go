package project

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (p *Project) Format(f *option.Format) string {
	return status.New(f).String(
		p.formatName(f),
		p.Path,
		p.formatConcern(f),
	).RawList(p).Format()
}
