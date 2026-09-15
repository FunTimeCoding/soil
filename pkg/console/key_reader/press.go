package key_reader

import "time"

func (r *Reader) Press(key rune) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	h, okay := r.handlers[key]

	if !okay {
		return
	}

	if _, held := r.pressed[key]; held {
		return
	}

	t := time.Now()
	r.pressed[key] = t

	if h.Press != nil {
		h.Press(key, t)
	}
}
