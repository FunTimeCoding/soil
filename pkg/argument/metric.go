package argument

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	metricConstant "github.com/funtimecoding/soil/pkg/metric/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func (i *Instance) Metric() {
	i.Integer(
		constant.MetricPort,
		environment.FallbackInteger(
			metricConstant.PortEnvironment,
			metricConstant.Port,
		),
		metricConstant.PortUsage,
	)
}
