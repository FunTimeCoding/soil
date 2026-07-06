package clean

import (
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/git/remote"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/hub"
)

func Hub(r *remote.Remote) {
	l := git.ParseLocator(r.Locator)

	if l == nil {
		system.Exitf(
			1,
			"could not parse remote locator: %s\n",
			r.Locator,
		)

		return
	}

	namespace, repository := git.ParseProject(l.Path)
	c := github.NewEnvironment()
	hub.Tag(c, namespace, repository)
	hub.Run(c, namespace, repository)
	hub.Image(c, namespace, repository)
}
