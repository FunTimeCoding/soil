package key_reader

import "time"

func (r *Reader) Release(key rune) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	h, okay := r.handlers[key]

	if !okay {
		return
	}

	t, held := r.pressed[key]

	if !held {
		return
	}

	delete(r.pressed, key)

	if h.Release != nil {
		h.Release(key, time.Since(t))
	}
}
