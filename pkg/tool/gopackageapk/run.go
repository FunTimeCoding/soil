package gopackageapk

import (
	"github.com/funtimecoding/soil/pkg/alpine/packager"
	"github.com/funtimecoding/soil/pkg/tool/gopackageapk/option"
)

func Run(o *option.Package) {
	p := packager.New(o.Executable, o.PackageVersion)
	p.CreateWorkspace()
	p.CopyBinary()
	p.WritePKGINFO(p.CreateArchive())
	p.CreateControlTar()
	p.ConcatenateTars()
	p.Cleanup()
}
