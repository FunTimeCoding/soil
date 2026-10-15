package gobuild

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/build"
	buildOption "github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobuild/option"
)

func Run(o *option.Build) error {
	names := o.Names

	if len(names) == 0 {
		names = programs()
	}

	for _, n := range names {
		console.Format("Build %s\n", n)
		p := buildOption.New()
		p.Name = n
		p.BuildTags = o.Tags
		p.Native = o.Native
		p.CopyToBin = o.Install
		p.Targets = o.Targets

		if o.Module != "" {
			p.Module = o.Module
			p.Version = o.Version

			if e := build.Module(p); e != nil {
				return e
			}

			continue
		}

		p.MainPath = o.Main

		if p.MainPath == "" {
			if p.MainPath = build.GuessMainPath(n); p.MainPath == "" {
				return fmt.Errorf("no %s/%s/main.go", constant.CommandPath, n)
			}
		}

		p.Output = o.Output
		build.Architectures(p)
	}

	return nil
}
