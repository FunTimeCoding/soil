package label_holder

import "time"

type Rule struct {
	Key           string
	Value         string
	Compatible    []string
	IdleThreshold time.Duration
	MinimumTurns  int
}
