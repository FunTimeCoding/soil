package client

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/alert"
	"time"
)

func toAlerts(responses []client.AlertsResponse) ([]*alert.Alert, error) {
	result := make([]*alert.Alert, 0, len(responses))

	for _, r := range responses {
		start, e := time.Parse(constant.DateFormat, r.Start)

		if e != nil {
			return nil, e
		}

		a := alert.New(uint(r.Identifier), r.Fingerprint, r.Name, start)
		a.Severity = r.Severity
		a.Summary = r.Summary

		if r.Labels != nil {
			a.Labels = r.Labels
		}

		if r.End != nil {
			end, f := time.Parse(constant.DateFormat, *r.End)

			if f != nil {
				return nil, f
			}

			a.End = &end
		}

		result = append(result, a)
	}

	return result, nil
}
