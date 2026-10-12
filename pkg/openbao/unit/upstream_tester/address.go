package upstream_tester

func (u *Tester) Address() string {
	return u.server.URL
}
