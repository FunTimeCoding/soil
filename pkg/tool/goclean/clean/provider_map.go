package clean

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/remote/provider_map"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/option"
	"github.com/funtimecoding/soil/pkg/web/host"
)

func providerMap(o *option.Clean) *provider_map.Map {
	result := provider_map.New()
	result.Add(o.GitLabHost, constant.GitLabProvider)

	if host.HasDot(o.GitLabHost) {
		result.Add(host.StripLeft(o.GitLabHost), constant.GitLabProvider)
	}

	return result
}
