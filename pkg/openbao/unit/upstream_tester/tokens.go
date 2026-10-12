package upstream_tester

func (u *Tester) Tokens() []string {
	u.lock.Lock()
	defer u.lock.Unlock()

	return u.tokens
}
