package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
)

func desiredStatus(
	guest string,
	current string,
	full bool,
) (string, bool) {
	if guest == constant.RunningStatus && current == netbox.StatusOffline {
		return netbox.StatusActive, true
	}

	if full && guest == constant.StoppedStatus &&
		current == netbox.StatusActive {
		return netbox.StatusOffline, true
	}

	return "", false
}
