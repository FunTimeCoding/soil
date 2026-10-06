package gopackage

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/build"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/gopackage/constant"
	"os"
)

func Main() {
	r := reporter.New(constant.Identity.Name()).Start()
	defer func() { r.RecoverFlush(recover()) }()
	a := argument.NewInstance(constant.Identity)
	a.Parse()
	var runs int

	for _, name := range build.OutputDirectories() {
		console.Format("Name: %s\n", name)
		outputDirectory := join.Relative(system.Temporary, name)
		console.Format("Output directory: %s\n", outputDirectory)

		for _, systemArchitecture := range build.SystemArchitectures() {
			if build.GuessBinaryPath(name, systemArchitecture) != "" {
				build.Archive(name, systemArchitecture)
				runs++
			}
		}
	}

	if runs == 0 {
		console.Line("No archive created")
		os.Exit(1)
	}
}
