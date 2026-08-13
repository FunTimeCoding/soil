package gobuild

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/build"
	buildConstant "github.com/funtimecoding/soil/pkg/build/constant"
	"github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/strings/split"
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
		argumentConstant.Main,
		"",
		"Path to main.go, defaults to cmd/$NAME/main.go",
	)
	a.String(
		argumentConstant.Output,
		"",
		"Output path, defaults to tmp/$NAME/$OS-$ARCH/$NAME",
	)
	a.String(argumentConstant.BuildTags, "", "Build tags")
	a.Boolean(buildConstant.CopyToBinFlag, false, "Copy to $HOME/bin")
	a.Boolean(systemConstant.LinuxAMD64, false, "Linux AMD64")
	a.Boolean(systemConstant.DarwinARM64, false, "Darwin ARM64")
	a.Boolean(systemConstant.DarwinAMD64, false, "Darwin AMD64")
	a.Boolean(buildConstant.Native, false, "Enable CGO")
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

	if name == argumentConstant.All {
		name = ""
	}

	if name == "" {
		for _, n := range system.Directories(
			join.Absolute(system.WorkDirectory(), systemConstant.CommandPath),
		) {
			if n == buildConstant.ExamplePath {
				continue
			}

			fmt.Printf("Build %s\n", n)
			buildNamed(a, n, linuxAMD64, darwinARM64, darwinAMD64)
		}

		return
	}

	if names := split.Comma(name); len(names) > 1 {
		for _, n := range names {
			fmt.Printf("Build %s\n", n)
			buildNamed(a, n, linuxAMD64, darwinARM64, darwinAMD64)
		}

		return
	}

	mainPath := a.GetString(argumentConstant.Main)

	if mainPath == "" {
		if mainPath = build.GuessMainPath(name); mainPath == "" {
			log.Panicf("could not find main.go for %s", name)
		}
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
