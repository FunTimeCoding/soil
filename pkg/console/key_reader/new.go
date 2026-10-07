package key_reader

import (
	"github.com/funtimecoding/soil/pkg/console/types/key_callback"
	"time"
)

func New() *Reader {
	return &Reader{
		handlers: make(map[rune]key_callback.Callback),
		pressed:  make(map[rune]time.Time),
	}
}
