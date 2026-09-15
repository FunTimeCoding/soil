package key_reader

import "time"

func New() *Reader {
	return &Reader{
		handlers: make(map[rune]Callback),
		pressed:  make(map[rune]time.Time),
	}
}
