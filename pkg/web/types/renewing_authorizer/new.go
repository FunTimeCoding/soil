package renewing_authorizer

func New(token string) *Authorizer {
	return &Authorizer{Token: token}
}
