package upstream_tester

func (u *Tester) Logins() int {
	u.lock.Lock()
	defer u.lock.Unlock()

	return u.logins
}
