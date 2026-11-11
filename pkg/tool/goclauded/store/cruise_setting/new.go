package cruise_setting

import "time"

func New(
	sessionIdentifier string,
	grace *time.Duration,
) *Setting {
	return &Setting{
		SessionIdentifier: sessionIdentifier,
		Grace:             grace,
	}
}
