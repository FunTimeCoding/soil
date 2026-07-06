package pod

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (p *Pod) Format(f *option.Format) string {
	return status.New(f).String(
		p.Name,
		p.Cluster,
	).RawList(p.Raw).Format()
}
