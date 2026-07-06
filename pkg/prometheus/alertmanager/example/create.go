package example

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/tool/common"
)

func Create() {
	c := common.Alertmanager()
	c.MustCreate(
		constant.HighMemoryUsage,
		"localhost:9090",
		"High memory usage detected",
		"Memory usage exceeded 90% for 5 minutes.",
		"memory_usage",
	)
}
