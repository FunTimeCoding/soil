package key_reader

import (
	"github.com/funtimecoding/soil/pkg/console/types/key_callback"
	"time"
)

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
	r.handlers[k] = key_callback.Callback{Press: press, Release: release}
}
