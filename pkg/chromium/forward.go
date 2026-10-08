package chromium

import (
	"github.com/funtimecoding/soil/pkg/chromium/event"
	"iter"
)

func forward[T any](
	events iter.Seq2[T, error],
	convert func(T) *event.Event,
	observe func(*event.Event),
) {
	for v, e := range events {
		if e != nil {
			return
		}

		observe(convert(v))
	}
}
