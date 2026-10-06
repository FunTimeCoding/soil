package unit

func (a *renewing) Renew() error {
	a.token = "fresh"
	a.renewed++

	return nil
}
