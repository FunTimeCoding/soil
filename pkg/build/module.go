package build

import (
	"github.com/funtimecoding/soil/pkg/build/constant"
	"github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system"
)

func Module(o *option.Build) error {
	v, e := ModuleOrigin(o.Module, o.Version)

	if e != nil {
		return e
	}

	console.Format("Module: %s@%s\n", o.Module, v.Version)
	console.Format("Hash: %s\n", v.Hash)
	directory := system.TemporaryDirectory(constant.ThrowawayPattern)
	defer system.Remove(directory)

	if f := throwaway(directory, o.Module, v.Version, o.Name); f != nil {
		return f
	}

	p := *o
	p.Version = v.Version
	p.Directory = directory
	Architectures(&p)

	return nil
}
