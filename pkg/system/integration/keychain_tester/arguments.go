package keychain_tester

func (k *Tester) Arguments() string {
	return k.read("arguments")
}
