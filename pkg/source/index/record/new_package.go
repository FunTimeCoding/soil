package record

func NewPackage(fingerprint string) *Package {
	return &Package{Fingerprint: fingerprint}
}
