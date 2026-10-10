package gobuild

import (
	"github.com/funtimecoding/soil/pkg/build"
	"github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/lint/installed"
	"github.com/funtimecoding/soil/pkg/system"
	"log"
	"runtime"
)

func rebuildStale() {
	for _, b := range installed.Outdated(git.FindDirectory()) {
		console.Format("Rebuild %s in %s\n", b.Name, b.Directory)
		system.ChangeDirectory(b.Directory)
		mainPath := build.GuessMainPath(b.Name)

		if mainPath == "" {
			log.Panicf("could not find main.go for %s", b.Name)
		}

		o := option.New()
		o.Name = b.Name
		o.MainPath = mainPath
		o.CopyToBin = true
		o.OperatingSystem = runtime.GOOS
		o.Architecture = runtime.GOARCH
		build.Go(o)
	}
}
