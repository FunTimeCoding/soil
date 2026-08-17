package lease

func (l *Lease) Holder() string {
	if !l.Held() {
		return ""
	}

	return *l.Raw.Spec.HolderIdentity
}
