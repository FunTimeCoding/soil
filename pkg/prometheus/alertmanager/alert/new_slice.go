package alert

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func NewSlice(
	v models.GettableAlerts,
	link *locator.Locator,
) []*Alert {
	var result []*Alert

	for _, a := range v {
		result = append(result, New(a, link))
	}

	return result
}
