package git

import (
	"github.com/funtimecoding/soil/pkg/git/remote"
	"github.com/funtimecoding/soil/pkg/git/remote/provider_map"
)

func Remote(
	path string,
	m *provider_map.Map,
	name string,
) *remote.Remote {
	for _, e := range Remotes(path, m) {
		if e.Name == name {
			return e
		}
	}

	return nil
}
