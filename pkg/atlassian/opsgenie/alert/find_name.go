package alert

import (
	"fmt"
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"strings"
)

func (a *Alert) findName() string {
	var message string
	var description string
	var details map[string]string

	if a.RawList != nil {
		// details and description not available
		message = a.RawList.Message
	}

	if a.RawDetail != nil {
		message = a.RawDetail.Message
		description = a.RawDetail.Description
		details = a.RawDetail.Details
	}

	if n, okay := details[constant.AlertnameLabel]; okay {
		return a.shortenAlert(n)
	}

	result := a.shortenAlert(message)

	if result == atlassian.OpsgenieUnknownName {
		result = a.nameFromDescription(description)

		if result == atlassian.OpsgenieUnknownName {
			if a.RawList != nil {
				fmt.Printf("Unknown name (simple): %+v\n", a.RawList)
			}

			if a.RawDetail != nil {
				fmt.Printf("Unknown name (full): %+v\n", a.RawDetail)
			}
		}
	}

	return strings.TrimSpace(result)
}
