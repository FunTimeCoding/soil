package kubernetes

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"

func New(c face.PodSource) *Collector {
	return &Collector{kubernetes: c}
}
