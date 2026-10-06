package connection

func (f *Failure) Fingerprint() []string {
	return []string{f.Kind, f.Host}
}
