package gobuild

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/build"
	"github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/gobuild/constant"
	"log"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.NewOptional(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.String(
		argument.Main,
		"",
		"Path to main.go, defaults to cmd/$NAME/main.go",
	)
	a.String(
		argument.Output,
		"",
		"Output path, defaults to tmp/$NAME/$OS-$ARCH/$NAME",
	)
	a.String(argument.BuildTags, "", "Build tags")
	a.Boolean(build.CopyToBinFlag, false, "Copy to $HOME/bin")
	a.Boolean(systemConstant.LinuxAMD64, false, "Linux AMD64")
	a.Boolean(systemConstant.DarwinARM64, false, "Darwin ARM64")
	a.Boolean(systemConstant.DarwinAMD64, false, "Darwin AMD64")
	a.Boolean(build.Native, false, "Enable CGO")
	a.Parse(version, gitHash, buildDate)
	linuxAMD64 := a.GetBoolean(systemConstant.LinuxAMD64)
	darwinARM64 := a.GetBoolean(systemConstant.DarwinARM64)
	darwinAMD64 := a.GetBoolean(systemConstant.DarwinAMD64)

	if !linuxAMD64 && !darwinARM64 && !darwinAMD64 {
		linuxAMD64 = true
		darwinARM64 = true
		darwinAMD64 = true
	}

	name := a.Argument(0)

	if name == "" {
		for _, n := range system.Directories(
			join.Absolute(
				system.WorkDirectory(),
				systemConstant.CommandPath,
			),
		) {
			if n == build.ExamplePath {
				continue
			}

			fmt.Printf("Build %s\n", n)
			mainPath := build.GuessMainPath(n)

			if mainPath == "" {
				log.Panicf("could not find main.go for %s", n)
			}

			o := option.New()
			o.Name = n
			o.MainPath = mainPath
			o.Output = a.GetString(argument.Output)
			o.BuildTags = a.GetString(argument.BuildTags)
			o.CopyToBin = a.GetBoolean(build.CopyToBinFlag)
			o.Native = a.GetBoolean(build.Native)
			o.LinuxAMD64 = linuxAMD64
			o.DarwinARM64 = darwinARM64
			o.DarwinAMD64 = darwinAMD64
			build.Architectures(o)
		}

		return
	}

	mainPath := a.GetString(argument.Main)

	if mainPath == "" {
		if mainPath = build.GuessMainPath(name); mainPath == "" {
			log.Panicf("could not find main.go for %s", name)
		}
	}

	o := option.New()
	o.Name = name
	o.MainPath = mainPath
	o.Output = a.GetString(argument.Output)
	o.BuildTags = a.GetString(argument.BuildTags)
	o.CopyToBin = a.GetBoolean(build.CopyToBinFlag)
	o.Native = a.GetBoolean(build.Native)
	o.LinuxAMD64 = linuxAMD64
	o.DarwinARM64 = darwinARM64
	o.DarwinAMD64 = darwinAMD64
	build.Architectures(o)
}
