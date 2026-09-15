package key_reader

func (r *Reader) Unregister(key rune) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.handlers, key)
	delete(r.pressed, key)
}
