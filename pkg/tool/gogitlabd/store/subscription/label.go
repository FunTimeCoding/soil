package subscription

func (s Subscription) Label() string {
	if s.Alias != "" {
		return s.Alias
	}

	return s.Reference
}
