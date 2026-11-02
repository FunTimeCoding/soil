package example

import (
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Create() {
	c := common.Alertmanager()
	defer c.Close()
	p := prometheus.NewEnvironment()
	defer p.Close()
	c.MustCreate(
		constant.HighMemoryUsage,
		"localhost:9090",
		"High memory usage detected",
		"Memory usage exceeded 90% for 5 minutes.",
		p.GraphLink("memory_usage"),
	)
}
