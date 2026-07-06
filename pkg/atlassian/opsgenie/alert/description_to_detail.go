package alert

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/detail"

func (a *Alert) descriptionToDetail() *detail.Prometheus {
	var result *detail.Prometheus

	if a.parseDescription != nil {
		result = a.parseDescription(a.Description)
	} else {
		result = detail.New()
	}

	return result
}
