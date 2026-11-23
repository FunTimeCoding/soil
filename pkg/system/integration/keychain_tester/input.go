package keychain_tester

func (k *Tester) Input() string {
	return k.read("input")
}
