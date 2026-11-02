package clean

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/git/remote"
	"github.com/funtimecoding/soil/pkg/gitlab"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/lab"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/option"
)

func Lab(
	o *option.Clean,
	origin *remote.Remote,
) {
	remoteLocator := git.ParseLocator(origin.Locator)

	if remoteLocator == nil {
		system.Exitf(1, "could not parse remote locator: %s\n", origin.Locator)

		return
	}

	namespace, repository := git.ParseProject(remoteLocator.Path)
	c := gitlab.New(
		o.GitLabHost,
		environment.Required(gitlabConstant.TokenEnvironment),
	)
	defer c.Close()
	x := context.Background()
	p := c.MustProjectByName(x, namespace, repository)

	if o.Verbose {
		console.Format("Project: %d %s %s\n", p.Identifier, p.Namespace, p.Name)
	}

	lab.Pipeline(x, o, c, p)
	lab.Registry(x, c, p)
	lab.Package(x, c, p)
	lab.Tag(x, c, p)
}
