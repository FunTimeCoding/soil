package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/build/origin"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
)

func Module(o *option.Build) {
	v := ModuleOrigin(o.Module, o.Version)

	if v == nil {
		errors.Warning("no origin for %s@%s", o.Module, o.Version)
		v = origin.New("", "")
	}

	directory := join.Absolute(
		system.WorkDirectory(),
		systemConstant.Temporary,
		o.Name,
	)
	system.EnsurePathExists(directory)
	console.Format("Module: %s@%s\n", o.Module, o.Version)
	console.Format("Hash: %s\n", v.Hash)
	console.Format("Output: %s\n", filepath.Join(directory, o.Name))
	r := run.New()
	r.Verbose = true
	r.Panic = false
	r.Environment(constant.NativeEnabled, stringConstant.BooleanFalse)
	r.Environment(constant.Binary, directory)
	r.Execute(
		constant.Go,
		constant.Install,
		constant.LinkerFlagsArgument,
		LinkerFlags(o.Version, ShortHash(v.Hash), Date()),
		constant.TagsArgument,
		Tags(o.BuildTags),
		fmt.Sprintf("%s/cmd/%s@%s", o.Module, o.Name, o.Version),
	)
}
