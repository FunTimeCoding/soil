package option

import "time"

type Monitor struct {
	Enabled             bool
	Interval            time.Duration
	Channels            []string
	Topics              []string
	MessageLimit        int
	NotificationChannel string
	Username            string
}
