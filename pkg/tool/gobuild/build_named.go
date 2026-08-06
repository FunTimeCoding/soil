package gobuild

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/build"
	buildConstant "github.com/funtimecoding/soil/pkg/build/constant"
	"github.com/funtimecoding/soil/pkg/build/option"
	"log"
)

func buildNamed(
	a *argument.Instance,
	name string,
	linuxAMD64 bool,
	darwinARM64 bool,
	darwinAMD64 bool,
) {
	mainPath := build.GuessMainPath(name)

	if mainPath == "" {
		log.Panicf("could not find main.go for %s", name)
	}

	o := option.New()
	o.Name = name
	o.MainPath = mainPath
	o.Output = a.GetString(argumentConstant.Output)
	o.BuildTags = a.GetString(argumentConstant.BuildTags)
	o.CopyToBin = a.GetBoolean(buildConstant.CopyToBinFlag)
	o.Native = a.GetBoolean(buildConstant.Native)
	o.LinuxAMD64 = linuxAMD64
	o.DarwinARM64 = darwinARM64
	o.DarwinAMD64 = darwinAMD64
	build.Architectures(o)
}
