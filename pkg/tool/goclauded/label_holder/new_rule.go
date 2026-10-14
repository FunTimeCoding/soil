package label_holder

import "time"

func NewRule(
	key string,
	value string,
	idleThreshold time.Duration,
	minimumTurns int,
) *Rule {
	return &Rule{
		Key:           key,
		Value:         value,
		IdleThreshold: idleThreshold,
		MinimumTurns:  minimumTurns,
	}
}
