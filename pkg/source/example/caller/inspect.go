package caller

import "github.com/funtimecoding/soil/pkg/source/example/probe"

func Inspect(p *probe.Probe) string {
	return p.Describe()
}
