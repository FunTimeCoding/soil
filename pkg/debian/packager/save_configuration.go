package packager

import (
	"github.com/funtimecoding/soil/pkg/debian"
	"github.com/funtimecoding/soil/pkg/system"
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
