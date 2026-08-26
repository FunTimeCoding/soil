package packager

import "github.com/funtimecoding/soil/pkg/debian"

func (p *Packager) Build() {
	debian.BuildPackage(p.PackageName)
}
