package mock_collector

import "github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"

type Collector struct {
	source     string
	placements []*placement.Placement
	failure    error
}
