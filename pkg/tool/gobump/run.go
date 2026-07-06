package gobump

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gobump/option"
)

func Run(o *option.Bump) {
	d := system.WorkDirectory()

	if !git.IsCleanCommand() {
		s := git.Status(d)

		if !git.IsClean(s, false) {
			system.Exitf(1, "not clean:\n%s\n", s.String())
		}
	}

	git.Fetch()
	var next *semver.Version

	if versions := git.Versions(d); len(versions) == 0 {
		next = semver.New("0.0.0")
	} else {
		next = semver.New(versions[len(versions)-1].String())
	}

	switch o.Increase {
	case patch:
		next.BumpPatch()
	case minor:
		next.BumpMinor()
	case major:
		next.BumpMajor()
	default:
		system.Exitf(1, "unexpected increase: %s\n", o.Increase)
	}

	nextString := key_value.Empty(constant.VersionPrefix, next.String())
	fmt.Printf("Tag: %s\n", nextString)
	git.Tag(nextString)
	git.Push()
}
