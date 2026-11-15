package result

func NewUsageTables(
	read []string,
	written []string,
) *UsageTables {
	return &UsageTables{Read: read, Written: written}
}
