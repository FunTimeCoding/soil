package lease

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/face"

func New(c face.LeaseSource) *Collector {
	return &Collector{opnsense: c}
}
