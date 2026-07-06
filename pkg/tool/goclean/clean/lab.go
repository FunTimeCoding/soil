package clean

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/git/remote"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
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
		system.Exitf(
			1,
			"could not parse remote locator: %s\n",
			origin.Locator,
		)

		return
	}

	namespace, repository := git.ParseProject(remoteLocator.Path)
	c := gitlab.New(
		o.GitLabHost,
		environment.Required(constant.TokenEnvironment),
	)
	p := c.MustProjectByName(namespace, repository)

	if o.Verbose {
		fmt.Printf(
			"Project: %d %s %s\n",
			p.Identifier,
			p.Namespace,
			p.Name,
		)
	}

	lab.Pipeline(o, c, p)
	lab.Registry(c, p)
	lab.Package(c, p)
	lab.Tag(c, p)
}
