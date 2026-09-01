package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	stringJoin "github.com/funtimecoding/soil/pkg/strings/join"
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

	console.Format("Name: %s\n", p.Name)
	console.Format("Output: %s\n", p.Output)
	path := filepath.Dir(p.Output)
	console.Format("Path: %s\n", path)
	system.EnsurePathExists(path)
	s := []string{
		constant.Go,
		constant.Build,
		constant.LinkerFlagsArgument,
		stringJoin.Space(
			constant.LinkerSetVariable,
			fmt.Sprintf("main.Version=%s", GitTag()),
			constant.LinkerSetVariable,
			fmt.Sprintf("main.GitHash=%s", GitHash()),
			constant.LinkerSetVariable,
			fmt.Sprintf("main.BuildDate=%s", Date()),
		),
	}

	if p.BuildTags != "" {
		s = append(s, constant.TagsArgument, p.BuildTags)
	}

	s = append(s, []string{constant.OutputArgument, p.Output, p.MainPath}...)
	r := run.New()
	r.Verbose = true
	r.Panic = false

	if !p.Native {
		r.Environment(constant.NativeEnabled, stringConstant.BooleanFalse)
	}

	r.Environment(constant.System, p.OperatingSystem)
	r.Environment(constant.Architecture, p.Architecture)

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
