package goatlas

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"

func sightingName(v client.Sighting) string {
	if v.Hostname == nil || *v.Hostname == "" || *v.Hostname == "*" {
		return v.HardwareAddress
	}

	return *v.Hostname
}
