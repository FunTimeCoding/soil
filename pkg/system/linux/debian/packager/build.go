package packager

import "github.com/funtimecoding/soil/pkg/system/linux/debian"

func (p *Packager) Build() {
	debian.BuildPackage(p.PackageName)
}
