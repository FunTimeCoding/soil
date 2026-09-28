package goatlas

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlas/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/generated/client"
)

func placeName(v client.Placement) string {
	if v.PlaceName == nil || *v.PlaceName == "" {
		return constant.Unplaced
	}

	return *v.PlaceName
}
