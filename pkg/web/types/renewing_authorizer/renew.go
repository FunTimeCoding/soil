package renewing_authorizer

func (a *Authorizer) Renew() error {
	a.Token = "fresh"
	a.Renewed++

	return nil
}
