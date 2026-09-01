package clean

import (
	"github.com/funtimecoding/soil/pkg/console"
	git "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/option"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"strings"
)

func Run(o *option.Clean) {
	o.GitLabHost = strings.TrimPrefix(o.GitLabHost, constant.SecurePrefix)
	m := providerMap(o)
	r := originRemote(o, m)

	switch r.Provider {
	case git.GitHubProvider:
		Hub(r)
	case git.GitLabProvider:
		Lab(o, r)
	case git.UnknownProvider:
		// TODO: Consider deleting tags except latest locally and pushing them to the server
		console.Line("Unknown provider, nothing to clean")
	}
}
