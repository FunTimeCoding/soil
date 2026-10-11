package packager

import "github.com/funtimecoding/soil/pkg/system"

func (p *Packager) CreateWorkspace() {
	system.Remove(p.WorkDirectory)
	system.MakeDirectory(p.ControlDirectory)
	system.MakeDirectory(p.ArchiveDirectory)
}
