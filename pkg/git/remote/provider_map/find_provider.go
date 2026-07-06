package provider_map

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"strings"
)

func (m *Map) FindProvider(host string) string {
	switch host {
	case constant.GitHubHost:
		return GitHubProvider
	case constant.GitLabHost:
		return GitLabProvider
	default:
		for knownHost, knownProvider := range m.Known {
			if host == knownHost {
				return knownProvider
			}

			// host may also have the port ":2222" in it
			if strings.HasPrefix(host, knownHost) {
				return knownProvider
			}
		}
	}

	return UnknownProvider
}
