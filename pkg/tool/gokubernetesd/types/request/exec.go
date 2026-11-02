package request

import "time"

type Exec struct {
	Name      string
	Namespace string
	Command   []string
	Container string
	Timeout   time.Duration
}
