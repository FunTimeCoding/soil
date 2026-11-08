package result

func NewCensusTables(
	read []string,
	written []string,
) *CensusTables {
	return &CensusTables{Read: read, Written: written}
}
