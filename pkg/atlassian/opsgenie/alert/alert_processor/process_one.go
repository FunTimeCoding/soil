package alert_processor

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"

func (p *Processor) ProcessOne(v *alert.Alert) *alert.Alert {
	var result *alert.Alert

	if p.enricher != nil {
		result = p.enricher.Run([]*alert.Alert{v})[0]
	}

	return result
}
