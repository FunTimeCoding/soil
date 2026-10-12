package clean

import (
	"github.com/funtimecoding/soil/pkg/console"
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/option"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"strings"
)

func Run(o *option.Clean) {
	o.GitLabHost = strings.TrimPrefix(o.GitLabHost, webConstant.SecurePrefix)
	m := providerMap(o)
	r := originRemote(o, m)

	switch r.Provider {
	case gitConstant.GitHubProvider:
		Hub(r)
	case gitConstant.GitLabProvider:
		Lab(o, r)
	case gitConstant.UnknownProvider:
		console.Line("Unknown provider, nothing to clean")
	}
}
