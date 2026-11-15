package failing_dialer

func (d *Dialer) Attempts() int {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	return d.attempts
}
