package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"strings"
)

func (a *Alert) findName() string {
	var message string
	var description string
	var details map[string]string

	if a.RawList != nil {
		message = a.RawList.Message
	}

	if a.RawDetail != nil {
		message = a.RawDetail.Message
		description = a.RawDetail.Description
		details = a.RawDetail.Details
	}

	if n, okay := details[prometheusConstant.AlertnameLabel]; okay {
		return a.shortenAlert(n)
	}

	result := a.shortenAlert(message)

	if result == constant.OpsgenieUnknownName {
		result = a.nameFromDescription(description)

		if result == constant.OpsgenieUnknownName {
			if a.RawList != nil {
				console.Format("Unknown name (simple): %+v\n", a.RawList)
			}

			if a.RawDetail != nil {
				console.Format("Unknown name (full): %+v\n", a.RawDetail)
			}
		}
	}

	return strings.TrimSpace(result)
}
