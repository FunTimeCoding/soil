package policy

func New(
	keyName string,
	domain string,
	types []string,
) *Policy {
	return &Policy{KeyName: keyName, Domain: domain, Types: types}
}
