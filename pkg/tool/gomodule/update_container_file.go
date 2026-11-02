package gomodule

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/github"
	githubConstant "github.com/funtimecoding/soil/pkg/github/constant"
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

	if token := os.Getenv(githubConstant.TokenEnvironment); token != "" {
		c := github.New(token)
		defer c.Close()

		if r := release.LatestCompatible(
			c.MustReleases(
				githubConstant.DelveNamespace,
				githubConstant.DelveRepository,
			),
			v,
		); r != nil {
			d = project.ReplaceDelveVersion(d, semverTrim.Trim(r.Name))
		}
	}

	system.SaveFile(file, d)
}
