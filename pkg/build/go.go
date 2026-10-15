package build

import (
	"github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/console"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
	"runtime"
)

func Go(o *option.Build) {
	p := *o

	if p.Output == "" {
		if p.Name == "" {
			panic("output empty and main not specified")
		}

		p.Output = join.Relative(
			systemConstant.Temporary,
			p.Name,
			SystemArchitecture(p.OperatingSystem, p.Architecture),
			p.Name,
		)
	}

	output, e := filepath.Abs(p.Output)
	errors.PanicOnError(e)
	console.Format("Name: %s\n", p.Name)
	console.Format("Output: %s\n", p.Output)
	path := filepath.Dir(p.Output)
	console.Format("Path: %s\n", path)
	system.EnsurePathExists(path)
	s := []string{
		soil.Go,
		soil.Build,
		soil.TagsArgument,
		Tags(p.BuildTags),
		soil.OutputArgument,
		output,
		target(&p),
	}
	r := run.New()
	r.Verbose = true
	r.Panic = false
	r.Directory = p.Directory

	if !p.Native {
		r.Environment(soil.NativeEnabled, stringsConstant.BooleanFalse)
	}

	r.Environment(soil.System, p.OperatingSystem)
	r.Environment(soil.Architecture, p.Architecture)

	if t := r.Start(s...); t != "" {
		console.Format("Output:\n%s", t)
	}

	errors.PanicOnError(r.Error)

	if p.CopyToBin &&
		runtime.GOOS == p.OperatingSystem &&
		runtime.GOARCH == p.Architecture {
		destination := join.Absolute(
			system.Home(),
			systemConstant.Binary,
			p.Name,
		)
		console.Format("Source: %s\n", p.Output)
		console.Format("Destination: %s\n", destination)
		system.ReplaceFile(p.Output, destination)
		system.Executable(destination)
	}

	console.Format("Size: %dM\n", system.FileSize(p.Output)/1024/1024)
}
