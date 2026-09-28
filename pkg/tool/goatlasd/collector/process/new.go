package process

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"

func New(
	c face.ProcessSource,
	place string,
) *Collector {
	return &Collector{process: c, place: place}
}
