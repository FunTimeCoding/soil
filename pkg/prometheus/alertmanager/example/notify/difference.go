package notify

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"

func difference(
	past []*alert.Alert,
	now []*alert.Alert,
) ([]*alert.Alert, []*alert.Alert, []*alert.Alert) {
	var add []*alert.Alert
	var stay []*alert.Alert
	var remove []*alert.Alert

	for _, n := range now {
		var found bool

		for _, o := range past {
			if n.MonitorIdentifier == o.MonitorIdentifier {
				found = true

				break
			}
		}

		if !found {
			add = append(add, n)
		}
	}

	for _, o := range past {
		var found bool

		for _, n := range now {
			if n.MonitorIdentifier == o.MonitorIdentifier {
				found = true

				break
			}
		}

		if found {
			stay = append(stay, o)
		} else {
			remove = append(remove, o)
		}
	}

	return add, stay, remove
}
