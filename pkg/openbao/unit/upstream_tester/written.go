package upstream_tester

func (u *Tester) Written() map[string]any {
	u.lock.Lock()
	defer u.lock.Unlock()

	return u.written
}
