package gomodule

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/release"
	"github.com/funtimecoding/soil/pkg/project"
	semverTrim "github.com/funtimecoding/soil/pkg/semver"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
)

func updateContainerFile(v *semver.Version) {
	file := containerFileName()

	if file == "" {
		return
	}

	d := project.ReplaceGoFromVersion(
		system.ReadFile(system.WorkDirectory(), file),
		v.String(),
	)

	if token := os.Getenv(constant.TokenEnvironment); token != "" {
		if r := release.LatestCompatible(
			github.New(token).MustReleases(
				constant.DelveNamespace,
				constant.DelveRepository,
			),
			v,
		); r != nil {
			d = project.ReplaceDelveVersion(d, semverTrim.Trim(r.Name))
		}
	}

	system.SaveFile(file, d)
}
