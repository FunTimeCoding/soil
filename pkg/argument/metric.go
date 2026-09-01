package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	metric "github.com/funtimecoding/soil/pkg/metric/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Metric() {
	i.Integer(
		constant.MetricPort,
		environment.FallbackInteger(metric.PortEnvironment, metric.Port),
		metric.PortUsage,
	)
}
