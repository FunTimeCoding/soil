package gopackage

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/build"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/gopackage/constant"
	"os"
)

func Main(
	version string,
	gitHash string,
	buildDate string,
) {
	r := reporter.New(constant.Identity.Name(), version).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse(version, gitHash, buildDate)
	var runs int

	for _, name := range build.OutputDirectories() {
		fmt.Printf("Name: %s\n", name)
		outputDirectory := join.Relative(system.Temporary, name)
		fmt.Printf("Output directory: %s\n", outputDirectory)

		for _, systemArchitecture := range build.SystemArchitectures() {
			if build.GuessBinaryPath(name, systemArchitecture) != "" {
				build.Archive(name, systemArchitecture)
				runs++
			}
		}
	}

	if runs == 0 {
		fmt.Println("No archive created")
		os.Exit(1)
	}
}
