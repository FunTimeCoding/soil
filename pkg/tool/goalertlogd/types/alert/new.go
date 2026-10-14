package alert

import "time"

func New(
	identifier uint,
	fingerprint string,
	name string,
	start time.Time,
) *Alert {
	return &Alert{
		Identifier:  identifier,
		Fingerprint: fingerprint,
		Name:        name,
		Labels:      map[string]string{},
		Start:       start,
	}
}
