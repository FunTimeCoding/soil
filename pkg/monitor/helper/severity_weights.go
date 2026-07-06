package helper

import "github.com/funtimecoding/soil/pkg/monitor/constant"

func SeverityWeights(
	critical float64,
	warning float64,
	information float64,
) map[constant.Severity]float64 {
	return map[constant.Severity]float64{
		constant.Critical:    critical,
		constant.Warning:     warning,
		constant.Information: information,
	}
}
