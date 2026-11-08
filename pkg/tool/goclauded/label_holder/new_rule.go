package label_holder

import "time"

func NewRule(
	key string,
	value string,
	compatible []string,
	idleThreshold time.Duration,
	minimumTurns int,
) *Rule {
	return &Rule{
		Key:           key,
		Value:         value,
		Compatible:    compatible,
		IdleThreshold: idleThreshold,
		MinimumTurns:  minimumTurns,
	}
}
