package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_result"
)

func UpgradeList(v *upgrade_result.Result) *server.UpgradeList {
	groups := make([]server.UpgradeGroup, len(v.Groups))

	for i, g := range v.Groups {
		groups[i] = UpgradeGroup(g)
	}

	return &server.UpgradeList{
		Groups:          groups,
		Hosts:           v.Hosts,
		Vulnerabilities: v.Vulnerabilities,
		Releases:        v.Releases,
	}
}
