package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_host"
)

func UpgradeHost(v *upgrade_host.Host) server.UpgradeHost {
	result := server.UpgradeHost{
		Name:      v.Name,
		Priority:  v.Priority,
		Installed: v.Installed,
		Candidate: v.Candidate,
	}

	if len(v.Labels) > 0 {
		result.Labels = &v.Labels
	}

	if len(v.Binaries) > 0 {
		result.Binaries = &v.Binaries
	}

	return result
}
