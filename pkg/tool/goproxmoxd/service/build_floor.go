package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/floor"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/sample"
)

func buildFloor(samples []*sample.Sample) *floor.Floor {
	result := floor.New()

	for _, v := range samples {
		if v.Failure == nil {
			addSample(result, v)
		}
	}

	sortFloor(result)

	return result
}
