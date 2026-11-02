package service

import (
	proxmoxConstant "github.com/funtimecoding/soil/pkg/proxmox/constant"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func desiredStatus(
	guest string,
	current string,
	full bool,
) (string, bool) {
	if guest == proxmoxConstant.RunningStatus && current == gonetboxdConstant.StatusOffline {
		return gonetboxdConstant.StatusActive, true
	}

	if full && guest == proxmoxConstant.StoppedStatus &&
		current == gonetboxdConstant.StatusActive {
		return gonetboxdConstant.StatusOffline, true
	}

	return "", false
}
