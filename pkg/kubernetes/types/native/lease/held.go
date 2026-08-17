package lease

func (l *Lease) Held() bool {
	return l.Raw.Spec.HolderIdentity != nil &&
		*l.Raw.Spec.HolderIdentity != ""
}
