package helper

import soil "github.com/funtimecoding/soil/pkg/constant"

func SeverityWeights(
	critical float64,
	warning float64,
	information float64,
) map[soil.Severity]float64 {
	return map[soil.Severity]float64{
		soil.Critical:    critical,
		soil.Warning:     warning,
		soil.Information: information,
	}
}
