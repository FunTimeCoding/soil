package provider_map

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"strings"
)

func (m *Map) FindProvider(host string) string {
	switch host {
	case constant.GitHubHost:
		return constant.GitHubProvider
	case constant.GitLabHost:
		return constant.GitLabProvider
	default:
		for knownHost, knownProvider := range m.Known {
			if host == knownHost {
				return knownProvider
			}

			if strings.HasPrefix(host, knownHost) {
				return knownProvider
			}
		}
	}

	return constant.UnknownProvider
}
