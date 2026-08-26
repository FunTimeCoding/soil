package packager

import "github.com/funtimecoding/soil/pkg/system"

func (p *Packager) Cleanup() {
	system.Remove(p.WorkDirectory)
}
