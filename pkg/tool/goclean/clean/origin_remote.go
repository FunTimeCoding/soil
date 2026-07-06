package clean

import (
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/git/remote"
	"github.com/funtimecoding/soil/pkg/git/remote/provider_map"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclean/clean/option"
)

func originRemote(
	o *option.Clean,
	m *provider_map.Map,
) *remote.Remote {
	result := git.Remote(system.WorkDirectory(), m, constant.OriginRemote)

	if result == nil {
		system.Exitf(
			1,
			"could not identify provider: %s\n",
			o.GitLabHost,
		)
		// make static analyzer happy
		result = remote.New("", "", "")
	}

	return result
}
