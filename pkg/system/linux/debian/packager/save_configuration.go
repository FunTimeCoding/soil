package packager

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/linux/debian"
)

func (p *Packager) SaveConfiguration() {
	system.MakeDirectory(p.ConfigurationRoot)
	system.SaveFile(
		p.ControlFile,
		debian.RenderControl(
			p.ExecutableName,
			p.Architecture,
			p.PackageVersion,
			p.MaintainerName,
			p.MaintainerMail,
		),
	)
}
