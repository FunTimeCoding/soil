package goatlas

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlas/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
)

func sightingPlace(v client.Sighting) string {
	if v.PlaceName == nil || *v.PlaceName == "" {
		return constant.Unclaimed
	}

	return *v.PlaceName
}
