package index

func NewPackageRecord(fingerprint string) *PackageRecord {
	return &PackageRecord{Fingerprint: fingerprint}
}
