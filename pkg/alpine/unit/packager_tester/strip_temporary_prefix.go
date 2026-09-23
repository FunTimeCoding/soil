package packager_tester

import (
	"github.com/funtimecoding/soil/pkg/alpine/packager"
	"github.com/funtimecoding/soil/pkg/system"
)

func StripTemporaryPrefix(p *packager.Packager) {
	d := "gopackageapk-goexample"
	p.WorkDirectory = system.StripUntilDirectory(p.WorkDirectory, d)
	p.ControlDirectory = system.StripUntilDirectory(p.ControlDirectory, d)
	p.ArchiveDirectory = system.StripUntilDirectory(p.ArchiveDirectory, d)
}
