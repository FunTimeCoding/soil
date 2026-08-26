package gopackagedeb

import (
	"github.com/funtimecoding/soil/pkg/debian/packager"
	"github.com/funtimecoding/soil/pkg/tool/gopackagedeb/option"
)

func Run(o *option.Package) {
	p := packager.New(
		o.Executable,
		o.PackageVersion,
		o.MaintainerName,
		o.MaintainerMail,
	)
	p.SaveConfiguration()

	if o.SystemdUnit {
		p.SaveUnit()
		p.SavePostInstall(o.UpgradeMode)
		p.SavePreRemove()
		p.SavePostRemove()
	}

	p.MoveBinary()
	p.Build()
}
