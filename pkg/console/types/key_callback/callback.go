package key_callback

import "time"

type Callback struct {
	Press func(
		k rune,
		t time.Time,
	)
	Release func(
		k rune,
		d time.Duration,
	)
}
