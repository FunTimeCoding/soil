package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"
)

func UpgradeGroup(v *upgrade_group.Group) server.UpgradeGroup {
	hosts := make([]server.UpgradeHost, len(v.Hosts))

	for i, h := range v.Hosts {
		hosts[i] = UpgradeHost(h)
	}

	return server.UpgradeGroup{
		Release:          v.Release,
		Source:           v.Source,
		Kernel:           v.Kernel,
		Hosts:            hosts,
		Vulnerabilities:  v.Vulnerabilities,
		Fixed:            v.Fixed,
		InstalledLowest:  v.InstalledLowest,
		InstalledHighest: v.InstalledHighest,
		Candidate:        v.Candidate,
		TopScore:         v.TopScore,
	}
}
