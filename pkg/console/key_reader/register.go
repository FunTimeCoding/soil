package key_reader

import "time"

func (r *Reader) Register(
	k rune,
	press func(
		rune,
		time.Time,
	),
	release func(
		rune,
		time.Duration,
	),
) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.handlers[k] = Callback{Press: press, Release: release}
}
