package constant

import (
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

var SeverityColor = map[string]string{
	prometheusConstant.CriticalSeverity: webConstant.BadColor,
	prometheusConstant.WarningSeverity:  webConstant.WarnColor,
}
