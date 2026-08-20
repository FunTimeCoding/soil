package subscription

func (s Subscription) Has(name string) bool {
	return s[name]
}
