package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/sample"
	"time"
)

func (s *Service) Refresh() []*sample.Sample {
	var result []*sample.Sample

	for _, i := range s.inventory.Instances {
		start := time.Now()
		v := sample.New(i.Name)
		v.Failure = s.attempt(
			i.Name,
			func(c face.ProxmoxClient) error { return s.sampleClient(c, v) },
		)
		v.Duration = time.Since(start)
		result = append(result, v)
	}

	s.storeFloor(buildFloor(result))

	return result
}
