package resolver

func (r *Resolver) BaseExists(value string) bool {
	return r.Exists(value) || r.SiblingExists(value)
}
