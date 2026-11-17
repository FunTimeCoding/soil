package base

import "time"

type setup struct {
	hold          time.Duration
	channelAlways bool
	sessionFiles  [][2]string
}
